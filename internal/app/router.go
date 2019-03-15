package app

import (
	"container/list"
	"net/http"
	"regexp"
	"strings"
	"sync"
)

type MatchFunc func(r *http.Request) bool

type Router struct {
	sync.RWMutex
	routes map[string]*Route
	list   *list.List
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]*Route),
		list:   list.New(),
	}
}

type Route struct {
	name   string
	match  MatchFunc
	handle http.HandlerFunc
}

func NewRoute(name string, handle http.HandlerFunc, match MatchFunc) *Route {
	return &Route{
		name,
		match,
		handle,
	}
}

func (rtr *Router) Dispatch(w http.ResponseWriter, r *http.Request) {
	if route := rtr.route(r); route != nil {
		route.handle(w, r)
		return
	}

	http.NotFound(w, r)
}

func (rtr *Router) route(r *http.Request) *Route {
	rtr.RLock()
	defer rtr.RUnlock()

	for e := rtr.list.Front(); e != nil; e = e.Next() {
		if route, ok := e.Value.(*Route); ok {
			if route.match(r) {
				return route
			}
		}
	}

	return nil
}

func (r *Router) AddRoutes(routes ...*Route) {
	r.Lock()
	defer r.Unlock()
	for _, route := range routes {
		r.addRoute(route)
	}
}

func (r *Router) AddRoute(route *Route) {
	r.Lock()
	defer r.Unlock()
	r.addRoute(route)
}

func (r *Router) addRoute(route *Route) {
	r.list.PushBack(route)
}

func PathIs(path string) MatchFunc {
	return func(r *http.Request) bool {
		return r.URL.Path == path
	}
}

func PathStartsWith(prefix string) MatchFunc {
	return func(r *http.Request) bool {
		return strings.HasPrefix(r.URL.Path, prefix)
	}
}

func PathRegex(expr string) MatchFunc {
	e := regexp.MustCompile(expr)
	return func(r *http.Request) bool {
		return e.MatchString(r.URL.Path)
	}
}

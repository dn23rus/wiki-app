package app

import (
	"context"
	"crypto/subtle"
	"encoding/base64"
	"log"
	"net/http"
	"strings"
	"time"
)

type Middleware func(next http.HandlerFunc) http.HandlerFunc

func LoggingMiddleware(stdlog *log.Logger) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				stdlog.Println(r.URL.Path, time.Since(start))
			}()
			next(w, r)
		}
	}
}

func MethodsMiddleware(methods ...string) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			for _, method := range methods {
				if r.Method == method {
					next(w, r)
					return
				}
			}
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
	}
}

func SlugMiddleware(prefix string) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			slug := ""
			if len(prefix) <= len(r.URL.Path) {
				slug = r.URL.Path[len(prefix):]
			}
			ctx := context.WithValue(r.Context(), "slug", slug)
			next(w, r.WithContext(ctx))
		}
	}
}

func HttpAuthMiddleware(user, password string) Middleware {

	var isAuthorized = func(u, p string) bool {
		return subtle.ConstantTimeCompare([]byte(u), []byte(user)) == 1 &&
			subtle.ConstantTimeCompare([]byte(p), []byte(password)) == 1
	}

	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

			s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
			if len(s) != 2 {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			b, err := base64.StdEncoding.DecodeString(s[1])
			if err != nil {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			pair := strings.SplitN(string(b), ":", 2)
			if len(pair) != 2 {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			if !isAuthorized(pair[0], pair[1]) {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			next(w, r)
		}
	}
}

func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

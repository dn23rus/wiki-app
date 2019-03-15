package app

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
	"time"
)

type Instance struct {
	httpServer     *http.Server
	stdlog, errlog *log.Logger
}

func NewInstance(stdlog, errlog *log.Logger) *Instance {
	s := &Instance{}
	s.stdlog = stdlog
	s.errlog = errlog
	return s
}

func (app *Instance) Start(config Config) error {

	dbGateway := NewDbGateway(config.Database)
	defer dbGateway.Close()

	tplRenderer := NewTplRenderer(config.TemplateDir)
	handlerCtx := NewContext(dbGateway, tplRenderer, app.errlog)

	router := NewRouter()
	router.AddRoutes(
		NewRoute("home",
			Chain(HomepageHandler(handlerCtx), MethodsMiddleware("HEAD", "GET")),
			PathIs("/")),
		NewRoute("view",
			Chain(ViewHandler(handlerCtx), SlugMiddleware("/view/"), MethodsMiddleware("HEAD", "GET")),
			PathStartsWith("/view/")),
		NewRoute("create",
			Chain(CreateHandler(handlerCtx), MethodsMiddleware("HEAD", "GET")),
			PathIs("/create"),
		),
		NewRoute("edit",
			Chain(EditHandler(handlerCtx), SlugMiddleware("/edit/"), MethodsMiddleware("HEAD", "GET")),
			PathStartsWith("/edit/")),
		NewRoute("save",
			Chain(SaveHandler(handlerCtx), SlugMiddleware("/save/"), MethodsMiddleware("POST")),
			PathStartsWith("/save"),
		),
	)

	handler := http.NewServeMux()
	handler.Handle("/", Chain(router.Dispatch, LoggingMiddleware(app.stdlog)))

	app.httpServer = &http.Server{
		Addr:    config.Server.Listen,
		Handler: handler,
	}

	app.stdlog.Println("Start listen on", config.Server.Listen)

	err := app.httpServer.ListenAndServe()
	if err != http.ErrServerClosed {
		app.Shutdown()
		//log.Fatal("Http Server stopped unexpected")
		return err
	} else {
		//log.Fatal("Http Server stopped")
		return nil
	}
}

func (app *Instance) Shutdown() {
	if app.httpServer != nil {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err := app.httpServer.Shutdown(ctx)
		if err != nil {
			app.Fatal(err)
		} else {
			app.httpServer = nil
		}
	}
}

func (app *Instance) Fatal(v ...interface{}) {
	app.stdlog.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

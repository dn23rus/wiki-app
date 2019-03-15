package main

import (
	"github.com/dn23rus/wiki-v2/internal/app"
	"log"
	"os"
	"os/signal"
)

var stdlog, errlog *log.Logger

func init() {
	stdlog = log.New(os.Stdout, "[Wiki app] ", log.Ldate|log.Ltime|log.LUTC)
	errlog = log.New(os.Stderr, "[Wiki app] ", log.Ldate|log.Ltime|log.LUTC)
}

func main() {

	//idle()

	config := app.NewConfig();
	if err := config.LoadFromFile("configs/main.json"); err != nil {
		errlog.Fatal((err))
	}

	application := app.NewInstance(stdlog, errlog)
	errlog.Fatal(application.Start(config));
}

func idle() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill)
	for {
		select {
		case <-interrupt:
			os.Exit(0)
		}
	}
}

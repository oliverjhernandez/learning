package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {

	addr := flag.String("addr", ":8000", "HTTP network address")

	flag.Parse()

	loggerHandler := slog.NewTextHandler(os.Stdout, nil)
	logger := slog.New(loggerHandler)

	app := &application{
		logger: logger,
	}

	logger.Info("starting server", "addr", *addr)
	err := http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

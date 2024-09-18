package main

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"casita/cmd/api"
	"casita/internal/db"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
)

func main() {
	logger := httplog.NewLogger("httplog", httplog.Options{
		JSON:             true,
		LogLevel:         slog.LevelDebug,
		Concise:          true,
		RequestHeaders:   true,
		MessageFieldName: "message",
		TimeFieldFormat:  time.RFC850,
		Tags: map[string]string{
			"version": "v1",
			"env":     "dev",
		},
		QuietDownRoutes: []string{
			"/",
			"/healthz",
		},
		QuietDownPeriod: 10 * time.Second,
		SourceFieldName: "source",
	})

	cfg := db.DBCfg{
		Env:  "dev",
		Port: 4000,
		DB: db.DBParams{
			Host:   "localhost",
			Port:   "5432",
			Name:   "casita",
			User:   "postgres",
			Passwd: "secret",
			SSL:    "disable",
		},
	}

	chi := chi.NewRouter()
	chi.Use(middleware.Timeout(60 * time.Second))
	chi.Use(httplog.RequestLogger(logger))

	api.InitializeRoutes(chi, cfg, logger)

	logger.Log(context.TODO(), 0, "starting server")

	if err := http.ListenAndServe(":3000", chi); err != nil {
		logger.Error(err.Error())
	}
}

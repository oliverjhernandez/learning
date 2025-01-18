package main

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"money_lovers/cmd/api"
	"money_lovers/internal/db"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
	"github.com/go-chi/render"
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
		QuietDownRoutes: []string{},
		QuietDownPeriod: 10 * time.Second,
		SourceFieldName: "source",
	})

	cfg := db.DBCfg{
		Env:  "dev",
		Port: 4000,
		DB: db.DBParams{
			Host:   "localhost",
			Port:   "5432",
			Name:   "money_lovers",
			User:   "postgres",
			Passwd: "secret",
			SSL:    "disable",
		},
	}

	chi := chi.NewRouter()
	chi.Use(httplog.RequestLogger(logger))
	chi.Use(middleware.Timeout(60 * time.Second))
	chi.Use(middleware.Recoverer)
	chi.Use(middleware.URLFormat)
	chi.Use(render.SetContentType(render.ContentTypeJSON))

	api.InitializeRoutes(chi, cfg, logger)
	logger.Info("Routes initialized")

	logger.Log(context.TODO(), 0, "starting server")

	if err := http.ListenAndServe(":4000", chi); err != nil {
		logger.Error(err.Error())
	}
}

package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"
	"snippetbox/internal/models"
	"text/template"

	"github.com/go-playground/form/v4"
	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	logger        *slog.Logger
	snippets      *models.SnippetModel
	templateCache map[string]*template.Template
	formDecoder   *form.Decoder
}

func main() {

	addr := flag.String("addr", ":8000", "HTTP network address")
	flag.String("dsn", os.Getenv("SNIPPETBOX_DB"), "MySQL data source name")

	flag.Parse()

	loggerHandler := slog.NewTextHandler(os.Stdout, nil)
	logger := slog.New(loggerHandler)

	db, err := openDB()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	tc, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	formDecoder := form.NewDecoder()

	app := &application{
		logger:        logger,
		snippets:      &models.SnippetModel{Db: db},
		templateCache: tc,
		formDecoder:   formDecoder,
	}

	logger.Info("starting server", "addr", *addr)
	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", os.Getenv("SNIPPETBOX_DB"))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

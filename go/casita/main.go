package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	"casita/cmd/api"
	"casita/internal/db"
	"casita/internal/jsonlog"

	"github.com/gofiber/fiber/v2"
)

type dbParams struct {
	host   string
	port   string
	name   string
	user   string
	passwd string
	ssl    string

	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

type config struct {
	env     string
	port    int
	db      dbParams
	limiter struct {
		rps     float64
		burst   int
		enabled bool
	}
}

type app struct {
	logger *jsonlog.Logger
	config config
}

func main() {
	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)
	cfg := config{
		env:  "dev",
		port: 4000,
		db: dbParams{
			host:   "localhost",
			port:   "5432",
			name:   "casita",
			user:   "postgres",
			passwd: "secret",
			ssl:    "disable",
		},
	}

	app := &app{
		logger: logger,
		config: cfg,
	}

	client, err := connectSQL(app.config.db)
	if err != nil {
		app.logger.PrintFatal(err, nil)
	}

	app.logger.PrintInfo("database connection pool stablished", nil)
	defer client.Close()

	stores := &db.Store{
		DB:           client,
		UserStore:    db.NewPGUserStore(client),
		TxnStore:     db.NewPGTransactionStore(client),
		CreditStore:  db.NewPGCreditStore(client),
		AccountStore: db.NewPGAccountStore(client),
	}

	fiberConfig := fiber.Config{ErrorHandler: api.ErrorHandler}
	fiberApp := fiber.New(fiberConfig)
	app.logger.PrintInfo("starting server", map[string]string{
		"address": app.config.db.host,
		"env":     app.config.env,
	})

	listenAddr := flag.String("listenAddr", ":4000", "The listen address of the API server")

	api.InitializeRoutes(stores, fiberApp)

	err = fiberApp.Listen(*listenAddr)
	if err != nil {
		app.logger.PrintFatal(err, nil)
	}
}

func connectSQL(dbParams dbParams) (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		dbParams.host,
		dbParams.port,
		dbParams.name,
		dbParams.user,
		dbParams.passwd,
		dbParams.ssl)

	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

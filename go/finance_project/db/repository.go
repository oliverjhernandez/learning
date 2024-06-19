package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

const (
	DBNAME  = "finance"
	DBURI   = "mongodb://localhost:27017"
	TDBNAME = "test-finance"
	TDBURI  = "mongodb://localhost:27017"
)

// TODO: Maybe create a general store that contains all of them
type Store struct {
	UserStore
}

type Dropper interface {
	Drop(ctx context.Context) error
}

// func NewStore(pgClient *sql.DB) (*Store, error) {
// 	return &Store{
// 		UserStore: NewPGUserStore(pgClient),
//	}, nil
// }

func ConnectSQL() *sql.DB {
	// TODO: set config getter
	dbHost := "localhost"
	dbPort := "5432"
	dbName := "your_db_name"
	dbUser := "your_user"
	dbPassword := ""
	dbSSL := "disable"

	// Connecto to DB
	log.Println("Connecting to dabase")
	connectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", dbHost, dbPort, dbName, dbUser, dbPassword, dbSSL)

	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}

	return db
}

type Params map[string]any

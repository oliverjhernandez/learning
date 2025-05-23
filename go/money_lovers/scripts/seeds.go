package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	"money_lovers/internal/db"
)

var (
	ctx          = context.Background()
	txStore      db.TransactionStore
	userStore    db.UserStore
	creditStore  db.CreditStore
	accountStore db.AccountStore
)

func tearDown(db *sql.DB) error {
	tables := []string{"users", "accounts", "transactions", "credits"}

	for _, table := range tables {
		query := "TRUNCATE TABLE " + table + " RESTART IDENTITY CASCADE"
		_, err := db.Exec(query)
		if err != nil {
			return err
		}
	}

	return nil
}

func seed() {
	passwd := "test123"

	params := db.CreateUser{
		FirstName: "Corina",
		LastName:  "Pulido",
		Email:     "corina@gmail.com",
		Passwd:    db.Password{Plaintext: &passwd},
	}

	userParams, err := db.NewUserFromParams(&params)
	if err != nil {
		log.Fatal(err)
	}

	user, err := userStore.InsertUser(ctx, nil, userParams)
	if err != nil {
		log.Printf("insert failed: %v", err)
		os.Exit(1)
	}

	accountParams := db.Account{
		Name:     "Main",
		Entity:   db.BANCOLOMBIA,
		Currency: db.COP,
		UserID:   user.ID,
	}

	account, err := accountStore.InsertAccount(ctx, nil, &accountParams)
	if err != nil {
		log.Printf("insert failed: %v", err)
		os.Exit(1)
	}

	transactions := []*db.Transaction{
		{
			Concept:     "Supermercado",
			Description: "Desayuno de la semana",
			Value:       320000,
			Date:        time.Now().Add(-7 * 24 * time.Hour),
			Relevance:   db.Essential,
			AccountID:   account.ID,
		},
		{
			Concept:     "Actividades",
			Description: "Natación",
			Value:       20000,
			Date:        time.Now().Add(-5 * 24 * time.Hour),
			Relevance:   db.Optional,
			AccountID:   account.ID,
		},
		{
			Concept:     "Ahorro",
			Description: "Mensual",
			Value:       100000,
			Date:        time.Now().Add(-3 * 24 * time.Hour),
			Relevance:   db.Important,
			AccountID:   account.ID,
		},
		{
			Concept:     "Ropa",
			Description: "Playa",
			Value:       5000,
			Date:        time.Now().Add(-9 * 24 * time.Hour),
			Relevance:   db.Optional,
			AccountID:   account.ID,
		},
	}

	for _, v := range transactions {
		_, err := txStore.InsertTransaction(ctx, nil, v)
		if err != nil {
			log.Printf("insert failed: %v", err)
			os.Exit(1)
		}
	}

	credits := []*db.Credit{
		{
			ClosingDay:   19,
			DueDay:       27,
			Identifier:   "q3hf489657439-42f89h5",
			Entity:       db.BANCOLOMBIA,
			Type:         db.HIPOTECARIO,
			Rate:         2.14,
			Total:        80000000,
			UserID:       user.ID,
			Installments: 12,
		},
		{
			ClosingDay:   16,
			DueDay:       25,
			Identifier:   "uofhehuiwrhurwfghvrw-34678",
			Entity:       db.AV_VILLAS,
			Type:         db.LIBRE,
			Rate:         1.8,
			Total:        10000000,
			UserID:       user.ID,
			Installments: 24,
		},
	}

	for _, v := range credits {
		_, err := creditStore.InsertCredit(ctx, nil, v)
		if err != nil {
			log.Printf("insert failed: %v", err)
			os.Exit(1)
		}
	}
}

func main() {
	dbParams := db.DBParams{
		Host:   "localhost",
		Port:   "5432",
		Name:   "money_lovers",
		User:   "postgres",
		Passwd: "secret",
		SSL:    "disable",
	}
	client, err := db.ConnectSQL(dbParams)
	if err != nil {
		log.Printf("initialization failed: %v", err)
		os.Exit(1)
	}
	defer client.Close()

	tearDown(client)

	txStore = db.NewPGTransactionStore(client)
	userStore = db.NewPGUserStore(client)
	creditStore = db.NewPGCreditStore(client)
	accountStore = db.NewPGAccountStore(client)

	seed()
}

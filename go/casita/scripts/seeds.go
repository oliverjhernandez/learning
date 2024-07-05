package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	"finance/db"
	"finance/models"
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
	params := models.CreateUser{
		FirstName: "Corina",
		LastName:  "Pulido",
		Email:     "corina@gmail.com",
		IsAdmin:   true,
		Passwd:    "test123",
	}

	user, err := models.NewUserFromParams(&params)
	if err != nil {
		log.Fatal(err)
	}

	userID, err := userStore.InsertUser(ctx, nil, user)
	if err != nil {
		log.Printf("insert failed: %v", err)
		os.Exit(1)
	}

	account := models.Account{
		Name:     "Main",
		Entity:   models.BANCOLOMBIA,
		Currency: models.COP,
		UserID:   userID,
	}

	accountID, err := accountStore.InsertAccount(ctx, nil, &account)
	if err != nil {
		log.Printf("insert failed: %v", err)
		os.Exit(1)
	}

	transactions := []*models.Transaction{
		{
			Concept:     "Supermercado",
			Description: "Desayuno de la semana",
			Value:       320000,
			Date:        time.Now().Add(-7 * 24 * time.Hour),
			Relevance:   models.Essential,
			AccountID:   accountID,
		},
		{
			Concept:     "Actividades",
			Description: "Natación",
			Value:       20000,
			Date:        time.Now().Add(-5 * 24 * time.Hour),
			Relevance:   models.Optional,
			AccountID:   accountID,
		},
		{
			Concept:     "Ahorro",
			Description: "Mensual",
			Value:       100000,
			Date:        time.Now().Add(-3 * 24 * time.Hour),
			Relevance:   models.Important,
			AccountID:   accountID,
		},
		{
			Concept:     "Ropa",
			Description: "Playa",
			Value:       5000,
			Date:        time.Now().Add(-9 * 24 * time.Hour),
			Relevance:   models.Optional,
			AccountID:   accountID,
		},
	}

	for _, v := range transactions {
		_, err := txStore.InsertTransaction(ctx, nil, v)
		if err != nil {
			log.Printf("insert failed: %v", err)
			os.Exit(1)
		}
	}

	credits := []*models.Credit{
		{
			ClosingDate:  time.Now(),
			DueDate:      time.Now().Add(time.Hour * 24 * 365),
			Identifier:   "q3hf489657439-42f89h5",
			Entity:       models.BANCOLOMBIA,
			Type:         models.HIPOTECARIO,
			Rate:         2.14,
			Total:        80000000,
			UserID:       userID,
			Installments: 12,
		},
		{
			ClosingDate:  time.Now().Add(time.Hour),
			DueDate:      time.Now().Add(time.Hour * 24 * 365 * 4),
			Identifier:   "uofhehuiwrhurwfghvrw-34678",
			Entity:       models.AV_VILLAS,
			Type:         models.LIBRE,
			Rate:         1.8,
			Total:        10000000,
			UserID:       userID,
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
	_, client, err := db.NewStore()
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

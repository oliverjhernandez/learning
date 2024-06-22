package main

import (
	"context"
	"log"
	"os"
	"time"

	"finance/db"
	"finance/models"
)

var (
	ctx         = context.Background()
	txStore     db.TransactionStore
	userStore   db.UserStore
	creditStore db.CreditStore
)

func seed() {
	transactions := []*models.Transaction{
		{
			Concept:     "Supermercado",
			Description: "Desayuno de la semana",
			Value:       320000,
			Date:        1716481693,
			Relevance:   models.Essential,
			Currency:    models.USD,
			Account:     models.SAVINGS,
		},
		{
			Concept:     "Actividades",
			Description: "Natación",
			Value:       20000,
			Date:        1746381693,
			Relevance:   models.Optional,
			Currency:    models.USD,
			Account:     models.SAVINGS,
		},
		{
			Concept:     "Ahorro",
			Description: "Mensual",
			Value:       100000,
			Date:        1746381693,
			Relevance:   models.Important,
			Currency:    models.USD,
			Account:     models.SAVINGS,
		},
		{
			Concept:     "Ropa",
			Description: "Playa",
			Value:       5000,
			Date:        1746381693,
			Relevance:   models.Optional,
			Currency:    models.USD,
			Account:     models.SAVINGS,
		},
	}

	users := []*models.User{
		{
			FirstName: "Corina",
			LastName:  "Pulido",
			Email:     "corinapulido@gmail.com",
			Passwd:    "123567qwerty",
		},
		{
			FirstName: "Oliver",
			LastName:  "Hernandez",
			Email:     "oliverjhernandez@gmail.com",
			Passwd:    "123567qwerty",
		},
		{
			FirstName: "Chuck",
			LastName:  "Norris",
			Email:     "notascratch@gmail.com",
			Passwd:    "123567qwerty",
		},
	}

	credits := []*models.Credit{
		{
			ClosingDate:  time.Now(),
			DueDate:      time.Now().Add(time.Hour * 24 * 365),
			Entity:       models.BANCOLOMBIA,
			Identifier:   "q3hf489657439-42f89h5",
			Type:         models.HIPOTECARIO,
			Rate:         2.14,
			Total:        80000000,
			Installments: 12,
		},
		{
			ClosingDate:  time.Now(),
			DueDate:      time.Now().Add(time.Hour * 24 * 365 * 4),
			Entity:       models.AV_VILLAS,
			Identifier:   "uofhehuiwrhurwfghvrw-34678",
			Type:         models.LIBRE_INVERSION,
			Rate:         1.8,
			Total:        10000000,
			Installments: 24,
		},
	}

	for _, v := range transactions {
		txStore.InsertTransaction(ctx, nil, v)
	}

	for _, v := range users {
		userStore.InsertUser(ctx, nil, v)
	}

	for _, v := range credits {
		creditStore.InsertCredit(ctx, nil, v)
	}
}

func main() {
	seed()
}

func init() {
	_, client, err := db.NewStore()
	if err != nil {
		log.Printf("initialization failed: %v", err)
		os.Exit(1)
	}
	defer client.Close()

	txStore = db.NewPGTransactionStore(client)
	userStore = db.NewPGUserStore(client)
	creditStore = db.NewPGCreditStore(client)
}

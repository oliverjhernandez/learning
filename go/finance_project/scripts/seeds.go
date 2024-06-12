package main

import (
	"context"
	"log"
	"time"

	"finance/db"
	"finance/types"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client      *mongo.Client
	ctx         = context.Background()
	txStore     db.TransactionStore
	userStore   db.UserStore
	creditStore db.CreditStore
)

func seed() {
	transactions := []*types.Transaction{
		{
			TransactionBase: types.TransactionBase{
				Concept:     "Supermercado",
				Description: "Desayuno de la semana",
				Value:       320000,
				Date:        1716481693,
				Relevance:   types.Essential,
				Currency:    types.USD,
				Account:     types.SAVINGS,
			},
		},
		{
			TransactionBase: types.TransactionBase{
				Concept:     "Actividades",
				Description: "Natación",
				Value:       20000,
				Date:        1746381693,
				Relevance:   types.Optional,
				Currency:    types.USD,
				Account:     types.SAVINGS,
			},
		},
		{
			TransactionBase: types.TransactionBase{
				Concept:     "Ahorro",
				Description: "Mensual",
				Value:       100000,
				Date:        1746381693,
				Relevance:   types.Important,
				Currency:    types.USD,
				Account:     types.SAVINGS,
			},
		},
		{
			TransactionBase: types.TransactionBase{
				Concept:     "Ropa",
				Description: "Playa",
				Value:       5000,
				Date:        1746381693,
				Relevance:   types.Optional,
				Currency:    types.USD,
				Account:     types.SAVINGS,
			},
		},
	}

	users := []*types.User{
		{
			UserBase: types.UserBase{
				FirstName: "Corina",
				LastName:  "Pulido",
				Email:     "corinapulido@gmail.com",
			},
			Passwd: "123567qwerty",
		},
		{
			UserBase: types.UserBase{
				FirstName: "Oliver",
				LastName:  "Hernandez",
				Email:     "oliverjhernandez@gmail.com",
			},
			Passwd: "123567qwerty",
		},
		{
			UserBase: types.UserBase{
				FirstName: "Chuck",
				LastName:  "Norris",
				Email:     "notascratch@gmail.com",
			},
			Passwd: "123567qwerty",
		},
	}

	credits := []*types.Credit{
		{
			CreditBase: types.CreditBase{
				ClosingDate:         time.Now(),
				DueDate:             time.Now().Add(time.Hour * 24 * 365),
				Entity:              types.BANCOLOMBIA,
				Identifier:          "q3hf489657439-42f89h5",
				Type:                types.HIPOTECARIO,
				Rate:                2.14,
				Total:               80000000,
				Number_Installments: 12,
			},
		},
		{
			CreditBase: types.CreditBase{
				ClosingDate:         time.Now(),
				DueDate:             time.Now().Add(time.Hour * 24 * 365 * 4),
				Entity:              types.AV_VILLAS,
				Identifier:          "uofhehuiwrhurwfghvrw-34678",
				Type:                types.LIBRE_INVERSION,
				Rate:                1.8,
				Total:               10000000,
				Number_Installments: 24,
			},
		},
	}

	for _, v := range transactions {
		txStore.InsertTransaction(ctx, v)
	}

	for _, v := range users {
		userStore.InsertUser(ctx, v)
	}

	for _, v := range credits {
		creditStore.InsertCredit(ctx, v)
	}
}

func main() {
	seed()
}

func init() {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Database(db.TDBNAME).Drop(ctx); err != nil {
		log.Fatal(err)
	}
	txStore = db.NewMongoTransactionStore(client, db.TDBNAME)
	userStore = db.NewMongoUserStore(client, db.TDBNAME)
	creditStore = db.NewMongoCreditStore(client, db.TDBNAME)
}

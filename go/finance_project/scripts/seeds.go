package main

import (
	"context"
	"log"

	"finance/db"
	"finance/types"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client    *mongo.Client
	ctx       = context.Background()
	txStore   db.TransactionStore
	userStore db.UserStore
)

func seed() {
	transactions := []*types.Transaction{
		{
			Concept:     "Supermercado",
			Description: "Desayuno de la semana",
			Value:       320000,
			Date:        1716481693,
			Relevance:   types.Essential,
			Currency:    types.USD,
			Account:     types.SAVINGS,
		},
		{
			Concept:     "Actividades",
			Description: "Natación",
			Value:       20000,
			Date:        1746381693,
			Relevance:   types.Optional,
			Currency:    types.USD,
			Account:     types.SAVINGS,
		},
		{
			Concept:     "Ahorro",
			Description: "Mensual",
			Value:       100000,
			Date:        1746381693,
			Relevance:   types.Important,
			Currency:    types.USD,
			Account:     types.SAVINGS,
		},
		{
			Concept:     "Ropa",
			Description: "Playa",
			Value:       5000,
			Date:        1746381693,
			Relevance:   types.Optional,
			Currency:    types.USD,
			Account:     types.SAVINGS,
		},
	}

	users := []*types.User{
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

	for _, v := range transactions {
		txStore.InsertTransaction(ctx, v)
	}

	for _, v := range users {
		userStore.InsertUser(ctx, v)
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

	if err := client.Database(db.DBNAME).Drop(ctx); err != nil {
		log.Fatal(err)
	}
	txStore = db.NewMongoTransactionStore(client, db.DBNAME)
	userStore = db.NewMongoUserStore(client, db.DBNAME)
}

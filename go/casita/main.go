package main

import (
	"flag"
	"fmt"

	"casita/cmd/api"
	"casita/internal/db"

	"github.com/gofiber/fiber/v2"
)

func main() {
	stores, client, err := db.NewStore()
	if err != nil {
		fmt.Printf("error loading db store")
	}
	defer client.Close()

	fiberConfig := fiber.Config{ErrorHandler: api.ErrorHandler}
	fiberApp := fiber.New(fiberConfig)
	listenAddr := flag.String("listenAddr", ":4000", "The listen address of the API server")

	api.Routes(stores, fiberApp)

	fiberApp.Listen(*listenAddr)
}

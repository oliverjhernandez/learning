package api

import (
	"log"
	"os"
	"testing"

	"finance/fixtures"

	"github.com/gofiber/fiber/v2"
)

var (
	ts          *fixtures.TestStore
	app         *fiber.App
	txHandler   *TransactionHandler
	userHandler *UserHandler
)

func TestMain(m *testing.M) {
	ts, err := fixtures.NewTestStore()
	if err != nil {
		log.Fatalf("setup failed: %v", err)
	}

	app = fiber.New()
	txHandler = NewTransactionHandler(ts.Store)
	userHandler = NewUserHandler(ts.Store)

	code := m.Run()

	// if err := ts.TearDown(); err != nil {
	// 	log.Fatalf("failed while tearing down: %v", err)
	// }

	os.Exit(code)
}

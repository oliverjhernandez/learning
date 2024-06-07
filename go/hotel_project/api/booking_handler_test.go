package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"hotel/db/fixtures"
	"hotel/types"

	"github.com/gofiber/fiber/v2"
)

func TestAdminGetBookings(t *testing.T) {
	db := setup(t)
	defer db.tearDown(t)

	var (
		adminUser      = fixtures.AddUser(db.Store, "Admin", "Admin", true)
		user           = fixtures.AddUser(db.Store, "James", "Foo", false)
		hotel          = fixtures.AddHotel(db.Store, "Stabby Stabby", "bermuda", 4, nil)
		room           = fixtures.AddRoom(db.Store, "small", true, 89, hotel.ID)
		from           = time.Now()
		till           = time.Now().Add(time.Duration(time.Hour * 24 * 5))
		app            = fiber.New()
		bookingHandler = NewBookingHandler(db.Store)
	)

	fixtures.AddBooking(db.Store, user.ID, room.ID, from, till)
	app.Group("/", JWTAuthentication(db.User))
	app.Get("/", bookingHandler.HandlerGetBookings)
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Add("X-Api-Token", CreateTokenFromUser(adminUser))

	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status code %d but got %s", resp.StatusCode, "200")
	}

	var bookings []*types.Booking
	if err := json.NewDecoder(resp.Body).Decode(&bookings); err != nil {
		t.Fatal(err)
	}

	if len(bookings) < 1 {
		t.Fatalf("expected at least 1 booking but got %d", len(bookings))
	}

	// test non admin cannot access the booking
	// nonAdminReq := httptest.NewRequest("GET", "/", nil)
	// nonAdminReq.Header.Add("X-Api-Token", CreateTokenFromUser(user))
	//
	// resp, err = app.Test(nonAdminReq)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// if resp.StatusCode == http.StatusOK {
	// 	t.Fatalf("expected not a 200 status code but got %d", resp.StatusCode)
	// }
}

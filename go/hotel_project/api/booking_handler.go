package api

import (
	"net/http"

	"hotel/db"
	"hotel/types"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type BookingHandler struct {
	store *db.Store
}

func NewBookingHandler(store *db.Store) *BookingHandler {
	return &BookingHandler{
		store,
	}
}

func (bh *BookingHandler) HandlerGetBookings(c *fiber.Ctx) error {
	filter := bson.M{}
	bookings, err := bh.store.Booking.GetBookings(c.Context(), filter)
	if err != nil {
		return err
	}
	return c.JSON(bookings)
}

func (bh *BookingHandler) HandlerGetBooking(c *fiber.Ctx) error {
	booking, err := bh.store.Booking.GetBookingByID(c.Context(), c.Params("id"))
	if err != nil {
		return err
	}

	user, ok := c.Context().UserValue("user").(*types.User)
	if !ok {
		return err
	}

	if booking.UserID != user.ID {
		return c.Status(http.StatusUnauthorized).JSON(genericResponse{
			Type: "error", Msg: "not authorized",
		})
	}

	return c.JSON(booking)
}

func (bh *BookingHandler) HandlerInsertBookings(c *fiber.Ctx) error {
	return nil
}

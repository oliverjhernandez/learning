package api

import (
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
		return ErrNotFound()
	}
	return c.JSON(bookings)
}

func (bh *BookingHandler) HandlerGetBooking(c *fiber.Ctx) error {
	booking, err := bh.store.Booking.GetBookingByID(c.Context(), c.Params("id"))
	if err != nil {
		return ErrNotFound()
	}

	user, ok := c.Context().UserValue("user").(*types.User)
	if !ok {
		return err
	}

	if booking.UserID != user.ID {
		return ErrUnauthorized()
	}

	return c.JSON(booking)
}

func (bh *BookingHandler) HandlerCancelBooking(c *fiber.Ctx) error {
	id := c.Params("id")
	booking, err := bh.store.Booking.GetBookingByID(c.Context(), id)
	if err != nil {
		return ErrNotFound()
	}

	user, err := getAuthenticatedUser(c)
	if err != nil {
		return err
	}

	if booking.UserID != user.ID {
		return ErrUnauthorized()
	}

	update := bson.M{
		"canceled": true,
	}
	if err := bh.store.Booking.UpdateBooking(c.Context(), c.Params("id"), update); err != nil {
		return err
	}
	return c.JSON(genericResponse{
		Type: "msg", Msg: "updated",
	})
}

func (bh *BookingHandler) HandlerInsertBookings(c *fiber.Ctx) error {
	return nil
}

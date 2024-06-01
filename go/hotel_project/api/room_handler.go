package api

import (
	"fmt"
	"net/http"
	"time"

	"hotel/db"
	"hotel/types"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RoomHandler struct {
	store *db.Store
}

type BookRoomParams struct {
	NumPersons int       `json:"numPersons"`
	FromDate   time.Time `json:"fromDate"`
	TillDate   time.Time `json:"tillDate"`
}

func (bp BookRoomParams) validate() error {
	now := time.Now()
	if now.After(bp.FromDate) || now.After(bp.TillDate) {
		return fmt.Errorf("cannot book a room in the past")
	}

	return nil
}

func NewBookRoomHandler(store *db.Store) *RoomHandler {
	return &RoomHandler{
		store,
	}
}

func (rh *RoomHandler) HandleGetRooms(c *fiber.Ctx) error {
	filter := bson.M{}
	rooms, err := rh.store.Room.GetRooms(c.Context(), filter)
	if err != nil {
		return err
	}

	return c.JSON(rooms)
}

func (rh *RoomHandler) HandleBookRoom(c *fiber.Ctx) error {
	var params BookRoomParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	if err := params.validate(); err != nil {
		return err
	}

	roomID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return err
	}

	user, ok := c.Context().Value("user").(*types.User)
	if !ok {
		return c.Status(http.StatusInternalServerError).JSON(genericResponse{
			Type: "error",
			Msg:  "internal server error",
		})
	}

	available, err := rh.isRoomAvailable(c, roomID, &params)
	if err != nil {
		return err
	}

	if !available {
		return c.Status(http.StatusBadRequest).JSON(genericResponse{
			Type: "error", Msg: fmt.Sprintf("room %s already booked", c.Params(("id"))),
		})
	}

	booking := types.Booking{
		UserID:     user.ID,
		RoomID:     roomID,
		FromDate:   params.FromDate,
		TillDate:   params.TillDate,
		NumPersons: params.NumPersons,
	}

	inserted, err := rh.store.Booking.InsertBooking(c.Context(), &booking)
	if err != nil {
		return err
	}

	return c.JSON(inserted)
}

func (rh *RoomHandler) isRoomAvailable(c *fiber.Ctx, roomID primitive.ObjectID, params *BookRoomParams) (bool, error) {
	filter := bson.M{
		"roomID": roomID,
		"fromDate": bson.M{
			"$gte": params.FromDate,
		},
		"tillDate": bson.M{
			"$lte": params.TillDate,
		},
	}
	bookings, err := rh.store.Booking.GetBookings(c.Context(), filter)
	if err != nil {
		return false, err
	}

	ok := len(bookings) == 0
	return ok, nil
}

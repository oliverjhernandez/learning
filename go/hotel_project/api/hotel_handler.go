package api

import (
	"hotel/db"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HotelHandler struct {
	store *db.Store
}

type HotelQueryParms struct {
	Rooms  bool
	Rating int
}

func (hs *HotelHandler) HandlerGetHotels(c *fiber.Ctx) error {
	var qparams HotelQueryParms
	if err := c.QueryParser(&qparams); err != nil {
		return err
	}

	hotels, err := hs.store.Hotel.GetHotels(c.Context(), nil)
	if err != nil {
		return err
	}

	return c.JSON(hotels)
}

func (hs *HotelHandler) HandlerGetHotel(c *fiber.Ctx) error {
	id := c.Params("id")

	hotel, err := hs.store.Hotel.GetHotelByID(c.Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(hotel)
}

func (hs *HotelHandler) HandlerGetRooms(c *fiber.Ctx) error {
	id := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"hotelID": oid}
	rooms, err := hs.store.Room.GetRooms(c.Context(), filter)
	if err != nil {
		return err
	}

	return c.JSON(rooms)
}

func NewHotelHandler(s *db.Store) *HotelHandler {
	return &HotelHandler{
		s,
	}
}

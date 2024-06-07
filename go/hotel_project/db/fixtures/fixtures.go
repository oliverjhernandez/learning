package fixtures

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"hotel/db"
	"hotel/types"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddUser(store *db.Store, fname, lname string, isAdmin bool) *types.User {
	params := types.CreateUserParams{
		FirstName: fname,
		LastName:  lname,
		Passwd:    "superarrecho", // fixed value
		Email:     fmt.Sprintf("%s@%s.com", strings.ToLower(fname), strings.ToLower(lname)),
	}
	user, err := types.NewUserFromParams(params)
	if err != nil {
		log.Fatal(err)
	}

	user.IsAdmin = isAdmin
	if _, err := store.User.InsertUser(context.TODO(), user); err != nil {
		log.Fatal(err)
	}

	return user
}

func AddRoom(store *db.Store, size string, seaside bool, price float64, hotelID primitive.ObjectID) *types.Room {
	room := &types.Room{
		Size:    size,
		Seaside: seaside,
		Price:   price,
		HotelID: hotelID,
	}

	newRoom, err := store.Room.InsertRoom(context.TODO(), room)
	if err != nil {
		return nil
	}

	return newRoom
}

func AddHotel(store *db.Store, name, location string, rating int, roomIDs []primitive.ObjectID) *types.Hotel {
	if roomIDs == nil {
		roomIDs = []primitive.ObjectID{}
	}

	hotel := types.Hotel{
		Name:     name,
		Location: location,
		Rooms:    roomIDs,
		Rating:   rating,
	}

	newHotel, err := store.Hotel.InsertHotel(context.TODO(), &hotel)
	if err != nil {
		log.Fatal(err)
	}

	return newHotel
}

func AddBooking(store *db.Store, userID, roomID primitive.ObjectID, from, till time.Time) *types.Booking {
	booking := &types.Booking{
		UserID:   userID,
		RoomID:   roomID,
		FromDate: from,
		TillDate: till,
	}
	newBooking, err := store.Booking.InsertBooking(context.TODO(), booking)
	if err != nil {
		return nil
	}
	return newBooking
}

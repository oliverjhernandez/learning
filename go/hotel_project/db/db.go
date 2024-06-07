package db

const (
	DBURI   = "mongodb://localhost:27017"
	DBNAME  = "hotel-reservation"
	TDBNAME = "test-hotel-reservation"
)

type Params map[string]any

type Store struct {
	User    UserStore
	Hotel   HotelStore
	Room    RoomStore
	Booking BookingStore
}

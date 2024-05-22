package db

const (
	DBURI   = "mongodb://localhost:27017"
	DBNAME  = "hotel-reservation"
	TDBNAME = "test-hotel-reservation"
)

// DEPRECATED
// func ToObjectID(id string) (primitive.ObjectID, error) {
// 	oid, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return oid, nil
// }

package db

// type MongoUserStore struct {
// 	client     *mongo.Client
// 	collection *mongo.Collection
// 	dbname     string
// }
//
// func (us *MongoUserStore) GetUsers(ctx context.Context) ([]*types.User, error) {
// 	var users []*types.User
// 	cur, err := us.collection.Find(ctx, bson.M{})
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	if err := cur.All(ctx, users); err != nil {
// 		return nil, err
// 	}
//
// 	return users, nil
// }
//
// func (us *MongoUserStore) GetUserByID(ctx context.Context, id string) (*types.User, error) {
// 	oid, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	filter := bson.M{"_id": oid}
//
// 	var user types.User
// 	if err := us.collection.FindOne(ctx, filter).Decode(&user); err != nil {
// 		return nil, err
// 	}
//
// 	return &user, nil
// }
//
// func (us *MongoUserStore) InsertUser(ctx context.Context, user *types.User) (*types.User, error) {
// 	res, err := us.collection.InsertOne(ctx, user)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	user.ID = res.InsertedID.(primitive.ObjectID)
// 	return user, nil
// }
//
// func (us *MongoUserStore) UpdateUser(ctx context.Context, filter Params, params *types.UpdateUserParams) error {
// 	// values := bson.D{
// 	// 	primitive.E{
// 	// 		Key: "$set", Value: params.ToBSON(),
// 	// 	},
// 	// }
//
// 	values := createUpdateDocument(params)
//
// 	res, err := us.collection.UpdateOne(ctx, filter, values)
// 	if err != nil {
// 		return err
// 	}
//
// 	fmt.Printf("Hello: %+v\n", res)
//
// 	return nil
// }
//
// func (us *MongoUserStore) DeleteUser(ctx context.Context, id string) error {
// 	oid, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return err
// 	}
//
// 	filter := bson.M{"_id": oid}
//
// 	_, err = us.collection.DeleteOne(ctx, filter)
// 	if err != nil {
// 		return err
// 	}
//
// 	return nil
// }
//
// func NewMongoUserStore(mc *mongo.Client, dbname string) *MongoUserStore {
// 	return &MongoUserStore{
// 		client:     mc,
// 		dbname:     dbname,
// 		collection: mc.Database(dbname).Collection(userCollection),
// 	}
// }
//
// func createUpdateDocument(fields *types.UpdateUserParams) bson.M {
// 	updates := bson.M{}
// 	setFields := bson.M{}
//
// 	userBase := fields.UserBase
// 	values := reflect.ValueOf(userBase)
// 	typeData := values.Type()
//
// 	for i := 0; i < typeData.NumField(); i++ {
// 		field := typeData.Field(i)
// 		val := values.Field(i)
// 		tag := strings.Split(field.Tag.Get("json"), ",")[0]
//
// 		if val.Len() != 0 {
// 			setFields[tag] = val
// 		}
//
// 	}
//
// 	updates["$set"] = setFields
// 	return updates
// }
//
// func isZeroType(value reflect.Value) bool {
// 	zero := reflect.Zero(value.Type()).Interface()
//
// 	switch value.Kind() {
// 	case reflect.Slice, reflect.Array, reflect.Chan, reflect.Map:
// 		return value.Len() == 0
// 	default:
// 		return reflect.DeepEqual(zero, value.Interface())
// 	}
// }

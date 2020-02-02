package mongo

import (
	"context"
	"log"

	"bitbucket.org/jawacompu10/addressbook/models"
	"bitbucket.org/jawacompu10/addressbook/repo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type dbAddress struct {
	ID      primitive.ObjectID `bson:"_id"`
	Address models.Address     `bson:"address"`
}

// Repo provides DB support for addressbook service with MongoDB
type Repo struct {
	client *mongo.Client
	ctx    context.Context
	dbInfo DBInfo
}

// DBInfo is the struct to store the connection details of the DB
type DBInfo struct {
	URL            string
	DBName         string
	CollectionName string
}

// GetAddressByID fetches and returns an address from the DB, filtered by the given ID
func (mr *Repo) GetAddressByID(addrID string) (models.Address, error) {
	dbAddr := &dbAddress{}
	addr := models.Address{}

	id, err := primitive.ObjectIDFromHex(addrID)
	if err != nil {
		err = &repo.InvalidIDError{ID: addrID}
		log.Println(err)
		return addr, err
	}
	filter := bson.M{
		"_id": id,
	}

	collection := mr.client.Database(mr.dbInfo.DBName).Collection(mr.dbInfo.CollectionName)
	err = collection.FindOne(mr.ctx, filter).Decode(&dbAddr)
	if err != nil {
		log.Println(err)
	}
	addr = dbAddr.toServiceAddress()
	return addr, err
}

// NewRepo creates and returns a new repo value to interact with DB
func NewRepo(dbInfo DBInfo) (*Repo, error) {
	ctx := context.Background()
	client, err := mongo.NewClient(options.Client().ApplyURI(dbInfo.URL))
	if err != nil {
		return nil, err
	}
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	log.Println("Connected to DB")

	return &Repo{
		client: client,
		ctx:    ctx,
		dbInfo: dbInfo,
	}, err

}

// GetUserAddresses returns addresses filtered by the given user ID
func (mr *Repo) GetUserAddresses(userID string) ([]models.Address, error) {
	addresses := make([]models.Address, 0)
	filter := bson.M{
		"address.userid": userID,
	}

	collection := mr.client.Database(mr.dbInfo.DBName).Collection(mr.dbInfo.CollectionName)
	cursor, err := collection.Find(mr.ctx, filter)
	if err != nil {
		return addresses, err
	}
	defer cursor.Close(mr.ctx)
	for cursor.Next(mr.ctx) {
		var addr dbAddress
		cursor.Decode(&addr)
		addresses = append(addresses, addr.toServiceAddress())
	}

	return addresses, err
}

// AddAddress adds an address to the DB
func (mr *Repo) AddAddress(addr models.Address) (models.Address, error) {
	collection := mr.client.Database(mr.dbInfo.DBName).Collection(mr.dbInfo.CollectionName)
	dbAddr := &dbAddress{}
	dbAddr.fromServiceAddress(addr)
	r, err := collection.InsertOne(mr.ctx, dbAddr)
	if err != nil {
		log.Println(err)
		return addr, err
	}

	id := r.InsertedID.(primitive.ObjectID).Hex()
	addr.ID = id
	return addr, nil
}

// UpdateAddress updates an existing address record in the DB
func (mr *Repo) UpdateAddress(addr models.Address) (models.Address, error) {
	dbAddr := &dbAddress{}

	id, err := primitive.ObjectIDFromHex(addr.ID)
	if err != nil {
		err = &repo.InvalidIDError{ID: addr.ID}
		log.Println(err)
		return addr, err
	}
	filter := bson.M{
		"_id": id,
	}

	returnOption := options.After
	collection := mr.client.Database(mr.dbInfo.DBName).Collection(mr.dbInfo.CollectionName)
	err = collection.FindOneAndUpdate(mr.ctx, filter, addr, &options.FindOneAndUpdateOptions{
		ReturnDocument: &returnOption,
	}).Decode(&dbAddr)
	if err != nil {
		log.Println(err)
	}
	addr = dbAddr.toServiceAddress()
	return addr, err
}

package storelayer

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo interface {
	Create() error
	Read() error
	Update() error
	Delete() error
	GetSingle() interface{}
}

// https://www.mongodb.com/docs/drivers/go/current/usage-examples/find/

func Create(collection string, app Mongo) (result *mongo.InsertOneResult, err error) {
	uri := os.Getenv("MONGODB_URI")
	db := os.Getenv("MONGODB_DATABASE")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable.")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		//panic(err)
		return
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			//panic(err)
			return
		}
	}()

	coll := client.Database(db).Collection(collection)

	result, err = coll.InsertOne(context.TODO(), app.GetSingle())
	if err != nil {
		return
	}
	return
}

func Read(collection string, app Mongo) (err error) {
	uri := os.Getenv("MONGODB_URI")
	db := os.Getenv("MONGODB_DATABASE")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable.")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		//panic(err)
		return
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			//panic(err)
			return
		}
	}()
	filter, err := bson.Marshal(app.GetSingle())
	if err != nil {
		return
	}
	coll := client.Database(db).Collection(collection)
	err = coll.FindOne(context.TODO(), filter).Decode(app.GetSingle())
	if err == mongo.ErrNoDocuments {
		fmt.Println("No document was found")
		return
	}
	if err != nil {
		return
	}
	return
}

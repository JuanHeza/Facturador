package storelayer

import (
	"context"
	"fmt"
	"log"
	"os"

	_"github.com/juanheza/facturador/helperlayer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo interface {
	Create() (*mongo.InsertOneResult, error)
	CreateMany() (*mongo.InsertManyResult, error)
	Read() error
	Update() error
	Delete() error
	GetSingle() interface{}
	GetList() interface{}
	GetListAsArray() []interface{}
	getCollection() string
}

// https://www.mongodb.com/docs/drivers/go/current/usage-examples/find/

func Create(app Mongo) (result *mongo.InsertOneResult, err error) {
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

	coll := client.Database(db).Collection(app.getCollection())

	result, err = coll.InsertOne(context.TODO(), app.GetSingle())
	if err != nil {
		return
	}
	return
}

func CreateMany(app Mongo) (result *mongo.InsertManyResult, err error) {
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

	coll := client.Database(db).Collection(app.getCollection())

	result, err = coll.InsertMany(context.TODO(), app.GetListAsArray())
	if err != nil {
		return
	}
	return
}

func Read(app Mongo) (err error) {
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
	coll := client.Database(db).Collection(app.getCollection())
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

func ReadMany(app Mongo) (err error) {
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
	coll := client.Database(db).Collection(app.getCollection())
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return
	}
	// end find

	if err = cursor.All(context.TODO(), app.GetList()); err != nil {
		return
	}

	return
}

func Update(app Mongo) (result *mongo.UpdateResult, err error) {
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
	coll := client.Database(db).Collection(app.getCollection())
	////////////////////////////////////
	id, _ := primitive.ObjectIDFromHex("5eb3d668b31de5d588f42a7a")
	filter := bson.D{{"_id", id}}
	////////////////////////////////////
	data, err := bson.Marshal(app.GetSingle())
	if err != nil {
		return
	}
	update := bson.D{{"$set", data}}
	result, err = coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return
	}
	return
}

func UpdateMany(app Mongo) (result *mongo.UpdateResult, err error) {
	uri := os.Getenv("MONGODB_URI")
	//db := os.Getenv("MONGODB_DATABASE")
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

	return
}

func Delete(app Mongo) (result *mongo.DeleteResult, err error) {
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
	coll := client.Database(db).Collection(app.getCollection())
	filter, err := bson.Marshal(app.GetSingle())
	if err != nil {
		return
	}
	//result, err = coll.DeleteOne(context.TODO(), filter)
	result, err = coll.DeleteMany(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	return
}

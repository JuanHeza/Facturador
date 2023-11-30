package storelayer

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo interface {
	Create() error
	Read() error
	Update() error
	Delete() error
	GetSingle() interface{}
	GetList() interface{}
	GetListAsArray() []interface{}
}

// https://www.mongodb.com/docs/drivers/go/current/usage-examples/find/

func Create(collection string, app Mongo) (result *mongo.InsertOneResult, results *mongo.InsertManyResult, err error) {
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
	results, err = coll.InsertMany(context.TODO(), app.GetListAsArray())
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

func ReadMany(collection string, app Mongo) (err error) {
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

func Update(collection string, app Mongo) (result *mongo.UpdateResult, err error) {
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

func UpdateMany(collection string, app Mongo) (result *mongo.UpdateResult, err error) {
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

func Delete(collection string, app Mongo) (result *mongo.DeleteResult, err error) {
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

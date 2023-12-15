package storelayer

import (
	"context"
	"fmt"
	"log"
	"os"

	_ "github.com/juanheza/facturador/helperlayer"
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
	GetSingle() Model
	GetList() (list *[]Model)
	GetListInterface() (list []interface{})
	getCollection() string
	getOptions() bson.M
	getProjection(projection string) bson.M
	SetList(cursor *mongo.Cursor) (err error)
}

type Model interface {
	ToJson() string
	ToFilter() (filter bson.M)
}

func Create(app Mongo) (result *mongo.InsertOneResult, err error) {
	uri := os.Getenv("MONGODB_URI")
	db := os.Getenv("MONGODB_DATABASE")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable.")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Println("Panic at conect@Create: ", err)
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Println("Panic at disconect@Create: ", err)
			panic(err)
		}
	}()

	coll := client.Database(db).Collection(app.getCollection())

	result, err = coll.InsertOne(context.TODO(), app.GetSingle())
	if err != nil {
		log.Println("Panic at InsertOne@Create: ", err)
		panic(err)
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
		log.Println("Panic at conect@CreateMany: ", err)
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Println("Panic at disconect@CreateMany: ", err)
			panic(err)
		}
	}()

	coll := client.Database(db).Collection(app.getCollection())

	result, err = coll.InsertMany(context.TODO(), app.GetListInterface())
	if err != nil {
		log.Println("Panic at InsertOne@CreateMany: ", err)
		panic(err)
	}
	return
}

////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////

func Read(app Mongo) (err error) {
	uri := os.Getenv("MONGODB_URI")
	db := os.Getenv("MONGODB_DATABASE")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable.")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Println("Panic at conect@Read: ", err)
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Println("Panic at disconect@Read: ", err)
			panic(err)
		}
	}()
	var filter bson.M = app.GetSingle().ToFilter()

	coll := client.Database(db).Collection(app.getCollection())
	err = coll.FindOne(context.TODO(), filter).Decode(app.GetSingle())
	if err == mongo.ErrNoDocuments {
		fmt.Println("No document was found")
		return
	}
	if err != nil {
		log.Println("Panic at FindOne@Read: ", err)
		panic(err)
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
		log.Println("Panic at conect@ReadMany: ", err)
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Println("Panic at disconect@ReadMany: ", err)
			panic(err)
		}
	}()
	//filter, err := bson.Marshal()
	if err != nil {
		return
	}
	coll := client.Database(db).Collection(app.getCollection())
	cursor, err := coll.Find(context.TODO(), bson.D{{}}, options.Find().SetProjection(app.getProjection("id")))
	if err != nil {
		log.Println("Panic at Find@ReadMany: ", err)
		panic(err)
	}
	// end find
	if err = app.SetList(cursor); err != nil {
		//if err = cursor.All(context.TODO(), &decode); err != nil {
		log.Println("Panic at cursor@ReadMany: ", err)
		panic(err)
	}
	return
}

func ReadOneCriteria(app Mongo, criteria bson.D, projection ...*options.FindOneOptions) (err error) {
	uri := os.Getenv("MONGODB_URI")
	db := os.Getenv("MONGODB_DATABASE")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable.")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Println("Panic at conect@ReadMany: ", err)
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Println("Panic at disconect@ReadMany: ", err)
			panic(err)
		}
	}()
	//filter, err := bson.Marshal()
	if err != nil {
		return
	}
	coll := client.Database(db).Collection(app.getCollection())
	err = coll.FindOne(context.TODO(), criteria, projection...).Decode(app.GetSingle())
	if err == mongo.ErrNoDocuments {
		fmt.Println("No document was found")
		return
	}
	if err != nil {
		log.Println("Panic at FindOne@Read: ", err)
		return
	}
	return
}

func ReadManyCriteria(app Mongo, criteria bson.D) (err error) {
	uri := os.Getenv("MONGODB_URI")
	db := os.Getenv("MONGODB_DATABASE")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable.")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Println("Panic at conect@ReadMany: ", err)
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Println("Panic at disconect@ReadMany: ", err)
			panic(err)
		}
	}()
	//filter, err := bson.Marshal()
	if err != nil {
		return
	}
	coll := client.Database(db).Collection(app.getCollection())
	cursor, err := coll.Find(context.TODO(), criteria, options.Find().SetProjection(app.getProjection("id")))
	if err != nil {
		log.Println("Panic at Find@ReadMany: ", err)
		panic(err)
	}
	// end find

	if err = cursor.All(context.TODO(), app.GetList()); err != nil {
		log.Println("Panic at cursor@ReadMany: ", err)
		panic(err)
	}
	return
}

////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////

func Update(app Mongo, id primitive.ObjectID) (result *mongo.UpdateResult, err error) {
	uri := os.Getenv("MONGODB_URI")
	db := os.Getenv("MONGODB_DATABASE")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable.")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Println("Panic at connect@Update: ", err)
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Println("Panic at disconnect@Update: ", err)
			panic(err)
		}
	}()
	coll := client.Database(db).Collection(app.getCollection())
	filter := bson.D{{Key: "_id", Value: id}}

	update := app.getOptions()
	result, err = coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println("Panic at updateOne@Update: ", err)
		panic(err)
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
		log.Println("Panic at connect@UpdateMany: ", err)
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Println("Panic at disconnect@UpdateMany: ", err)
			panic(err)
		}
	}()

	return
}

////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////

func Delete(app Mongo) (result *mongo.DeleteResult, err error) {
	uri := os.Getenv("MONGODB_URI")
	db := os.Getenv("MONGODB_DATABASE")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable.")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Println("Panic at connect@Delete: ", err)
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Println("Panic at disconnect@Delete: ", err)
			panic(err)
		}
	}()
	coll := client.Database(db).Collection(app.getCollection())
	filter, err := bson.Marshal(app.GetSingle())
	if err != nil {
		log.Println("Panic at marshal@Delete: ", err)
		panic(err)
	}
	//result, err = coll.DeleteOne(context.TODO(), filter)
	result, err = coll.DeleteMany(context.TODO(), filter)
	log.Println("Panic at deleteMany@Delete: ", err)
	if err != nil {
		panic(err)
	}
	return
}

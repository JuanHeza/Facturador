package storelayer

import (
	"log"

	"github.com/juanheza/facturador/helperlayer"
	"github.com/juanheza/facturador/modellayer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStore struct {
	Single  *modellayer.User   `json:"-"`
	List    []*modellayer.User `json:"-"`
	Options bson.M
}

func (us *UserStore) Create() (res *mongo.InsertOneResult, err error) {
	us.Single.Control.Init(1)
	res, err = Create(us)
	return
}

func (us *UserStore) CreateMany() (res *mongo.InsertManyResult, err error) {
	us.Single.Control.Init(1)
	res, err = CreateMany(us)
	return
}

func (us *UserStore) Read() (err error) {
	err = Read(us)
	log.Println(us)
	return
}

func (us *UserStore) Update() (err error) {
	result, err := Update(us, us.Single.UsuarioID)
	log.Println(result)
	return
}

func (us *UserStore) Delete() (err error) {
	result, err := Delete(us)
	log.Println(result)
	return
}

func (us *UserStore) GetSingle() (single Model) {
	single = us.Single
	return
}

func (us *UserStore) GetList() (list *[]Model) {
    aux := []Model{}
	for _, d := range us.List {
        aux = append(aux, d)
    }
    list = &aux
	return
}

func (us *UserStore) GetListInterface() (list []interface{}){
    var models = us.GetList()
    for _, model := range *models {
        list = append(list, model)
    }
    return
}

func (us *UserStore) getCollection() string {
	return helperlayer.User.ToString()
}

func (us *UserStore) getOptions() bson.M {
	return us.Options
}

func (us *UserStore) getProjection(projection string) bson.M{
    projectionCatalog := map[string]bson.M{
        "id": bson.M{
            "_id": 1,
        },
    }
    return projectionCatalog[projection]
}

package storelayer

import (
	"log"

	"github.com/juanheza/facturador/helperlayer"
	"github.com/juanheza/facturador/modellayer"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStore struct {
    Single *modellayer.User   `json:"-"`
    List   []*modellayer.User `json:"-"`
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
    result, err := Update(us)
    log.Println(result)
    return
}

func (us *UserStore) Delete() (err error) {
    result, err := Delete(us)
    log.Println(result)
    return
}

func (us *UserStore) GetSingle() (single interface{}) {
    single = us.Single
    return
}

func (us *UserStore) GetList() (list interface{}) {
    list = us.List
    return
}

func (us *UserStore) GetListAsArray() (list []interface{}) {
    for _, element := range us.List {
        list = append(list, element)
    }
    return
}

func (us *UserStore) getCollection() string{
    return helperlayer.User.ToString()
}


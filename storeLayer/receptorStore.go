package storelayer

import (
    "log"

    "github.com/juanheza/facturador/helperlayer"
    "github.com/juanheza/facturador/modellayer"
    "go.mongodb.org/mongo-driver/mongo"
)

type ReceptorStore struct {
    Single *modellayer.Receptor   `json:"-"`
    List   []*modellayer.Receptor `json:"-"`
}

func (rc *ReceptorStore) Create() (res *mongo.InsertOneResult, err error) {
    rc.Single.Control.Init(1)
    res, err = Create(rc)
    return
}

func (rc *ReceptorStore) CreateMany() (res *mongo.InsertManyResult, err error) {
    rc.Single.Control.Init(1)
    res, err = CreateMany(rc)
    return
}

func (rc *ReceptorStore) Read() (err error) {
    err = Read(rc)
    log.Println(rc)
    return
}

func (rc *ReceptorStore) Update() (err error) {
    result, err := Update(rc)
    log.Println(result)
    return
}

func (rc *ReceptorStore) Delete() (err error) {
    result, err := Delete(rc)
    log.Println(result)
    return
}

func (rc *ReceptorStore) GetSingle() (single interface{}) {
    single = rc.Single
    return
}

func (rc *ReceptorStore) GetList() (list interface{}) {
    list = rc.List
    return
}

func (rc *ReceptorStore) GetListAsArray() (list []interface{}) {
    for _, element := range rc.List {
        list = append(list, element)
    }
    return
}

func (rc *ReceptorStore) getCollection() string{
    return helperlayer.Receptor.ToString()
}


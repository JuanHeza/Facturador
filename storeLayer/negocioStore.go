package storelayer

import (
	"log"

	"github.com/juanheza/facturador/helperlayer"
	"github.com/juanheza/facturador/modellayer"
	"go.mongodb.org/mongo-driver/mongo"
)

type NegocioStore struct {
	Single *modellayer.Negocio   `json:"-"`
	List   []*modellayer.Negocio `json:"-"`
}

func (ns *NegocioStore) Create() (res *mongo.InsertOneResult, err error) {
    ns.Single.Control.Init(1)
    res, err = Create(ns)
    return
}
func (ns *NegocioStore) CreateMany() (res *mongo.InsertManyResult, err error) {
    ns.Single.Control.Init(1)
    res, err = CreateMany(ns)
    return
}

func (ns *NegocioStore) Read() (err error) {
	err = Read(ns)
	log.Println(ns)
	return
}
func (ns *NegocioStore) Update() (err error) {
	result, err := Update(ns)
	log.Println(result)
	return
}
func (ns *NegocioStore) Delete() (err error) {
	result, err := Delete(ns)
	log.Println(result)
	return
}
func (ns *NegocioStore) GetSingle() (single interface{}) {
	single = ns.Single
	return
}
func (ns *NegocioStore) GetList() (list interface{}) {
    list = ns.List
    return
}
func (ns *NegocioStore) GetListAsArray() (list []interface{}) {
    for _, element := range ns.List {
        list = append(list, element)
    }
    return
}

func (ns *NegocioStore) getCollection() string{
    return helperlayer.Negocio.ToString()
}

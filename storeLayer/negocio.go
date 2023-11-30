package storelayer

import (
	"log"

	"github.com/juanheza/facturador/modellayer"
)

type NegocioStore struct {
	Single *modellayer.Negocio   `json:"-"`
	List   []*modellayer.Negocio `json:"-"`
}

var (
	collection string = "negocio"
)

func (ns *NegocioStore) Create() (err error) {
    ns.Single.Control.Init(1)
	one, many, err := Create(collection, ns)
    log.Println(one)
    log.Println(many)
	return
}

func (ns *NegocioStore) Read() (err error) {
	err = Read(collection, ns)
	log.Println(ns)
	return
}
func (ns *NegocioStore) Update() (err error) {
	result, err := Update(collection, ns)
	log.Println(result)
	return
}
func (ns *NegocioStore) Delete() (err error) {
	result, err := Delete(collection, ns)
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

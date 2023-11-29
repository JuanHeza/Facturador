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
	result, err := Create(collection, ns)
	log.Println(result)
	return
}

func (ns *NegocioStore) Read() (err error) {
	err = Read(collection, ns)
	log.Println(ns)
	return
}
func (ns *NegocioStore) Update() (err error) {
	return
}
func (ns *NegocioStore) Delete() (err error) {
	return
}
func (ns *NegocioStore) GetSingle() (single interface{}) {
	single = ns.Single
	return
}

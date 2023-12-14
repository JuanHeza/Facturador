package storelayer

import (
	"fmt"
	"log"

	"github.com/juanheza/facturador/helperlayer"
	"github.com/juanheza/facturador/modellayer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type NegocioStore struct {
	Single *modellayer.Negocio   `json:"-"`
	List   []*modellayer.Negocio `json:"-"`
	Options bson.M
}
func NewNegocioStore() NegocioStore{
    return NegocioStore{
        Single: modellayer.NewNegocio(),
        List: []*modellayer.Negocio{},
        Options: bson.M{},
    }
}
func (ns *NegocioStore) Create() (res *mongo.InsertOneResult, err error) {
	ns.Single.Control.Init(1)
	res, err = Create(ns)
    log.Println(err)
	return
}
func (ns *NegocioStore) CreateMany() (res *mongo.InsertManyResult, err error) {
	ns.Single.Control.Init(1)
	res, err = CreateMany(ns)
    log.Println(err)
	return
}

func (ns *NegocioStore) Read() (err error) {
    fmt.Println("READING")
	err = Read(ns)
	log.Println(ns)
    log.Println(err)
	return
}
func (ns *NegocioStore) ReadMany() (err error) {
    fmt.Println("READING MANY")
	err = ReadMany(ns)
	log.Println(ns)
    log.Println(err)
	return
}
func (ns *NegocioStore) Update() (err error) {
	result, err := Update(ns, ns.Single.NegocioID)
	log.Println(result)
    log.Println(err)
	return
}
func (ns *NegocioStore) Delete() (err error) {
	result, err := Delete(ns)
	log.Println(result)
    log.Println(err)
	return
}
func (ns *NegocioStore) GetSingle() (single Model) {
	single = ns.Single
	return
}
func (ns *NegocioStore) GetList() (list *[]Model) {
    aux := []Model{}
    for _, d := range ns.List {
        aux = append(aux, d)
    }
    list = &aux
	return
}
func (ns *NegocioStore) GetListInterface() (list []interface{}){
    var models = ns.GetList()
    for _, model := range *models {
        list = append(list, &model)
    }
    return
}
func (ns *NegocioStore) getCollection() string {
	return helperlayer.Negocio.ToString()
}
func (ns *NegocioStore) getOptions() bson.M{
	return ns.Options
}
func (ns *NegocioStore) getProjection(projection string) bson.M{
    projectionCatalog := map[string]bson.M{
        "id": bson.M{
            "_id": 1,
            "colorSecundario": 1,
            "colorPrincipal": 1,
            "logo": 1,
            "clave": 1,
            
        },
        "expandesd": bson.M{
            "_id": 1,
        },
    }
	return projectionCatalog[projection]
}
func (ns *NegocioStore) UpdateTokenAdmin(user *modellayer.User){
	ns.Options = bson.M{
		"$set": bson.M{
			"bearerToken": ns.Single.BearerToken,
            "owner": user,
		},
        /*
		"$push": bson.M{
			"usuarios": user,
		},
  */
	}
}

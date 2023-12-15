package storelayer

import (
	"log"

	"github.com/juanheza/facturador/helperlayer"
	"github.com/juanheza/facturador/modellayer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ReceptorStore struct {
	Single  *modellayer.Receptor   `json:"-"`
	List    []*modellayer.Receptor `json:"-"`
	Options bson.M
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
	result, err := Update(rc, rc.Single.ReceptorID)
	log.Println(result)
	return
}

func (rc *ReceptorStore) Delete() (err error) {
	result, err := Delete(rc)
	log.Println(result)
	return
}

func (rc *ReceptorStore) GetSingle() (single Model) {
	single = rc.Single
	return
}

func (rc *ReceptorStore) GetList() (list *[]Model) {
    aux := []Model{}
	for _, d := range rc.List {
        aux = append(aux, d)
    }
    list = &aux
	return
}
func (rc *ReceptorStore) GetListInterface() (list []interface{}){
    var models = rc.GetList()
    for _, model := range *models {
        list = append(list, model)
    }
    return
}


func (rc *ReceptorStore) getCollection() string {
	return helperlayer.Receptor.ToString()
}
func (rc *ReceptorStore) getOptions() bson.M {
	return rc.Options
}
func (rc *ReceptorStore) SetList(list []interface{}) {
	dst := make([]*modellayer.Receptor, len(list))
	for i := range list {
		dst[i] = list[i].(*modellayer.Receptor)
	}
	rc.List = dst
}
func (rc *ReceptorStore) getProjection(projection string) bson.M{
    projectionCatalog := map[string]bson.M{
        "id": bson.M{
            "_id": 1,
        },
    }
    return projectionCatalog[projection]
}
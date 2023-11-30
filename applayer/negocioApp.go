package applayer

import (
	"log"

	"github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/juanheza/facturador/modellayer"
	storelayer "github.com/juanheza/facturador/storeLayer"
)

type NegocioApp struct {
	Response modellayer.Response
	store    storelayer.NegocioStore
}
var negocioApp NegocioApp

func (ng *NegocioApp) Create(context *gin.Context) {
	// Call BindJSON to bind the received JSON
	negocioApp.decode(context)
    res, err := negocioApp.store.Create(); 
	if err != nil{
        return 
    }
    userApp := &UserApp{}

    if id, ok := res.InsertedID.(primitive.ObjectID); ok {

    userApp.CreateAdmin(context, id)
    }
	return
}
func (ng *NegocioApp) Read(context *gin.Context) {
	negocioApp.decode(context)
	negocioApp.store.Read()
	return
}
func (ng *NegocioApp) Update(context *gin.Context) {
	negocioApp.decode(context)
	negocioApp.store.Update()
	return
}
func (ng *NegocioApp) Delete(context *gin.Context) {
	negocioApp.decode(context)
	negocioApp.store.Delete()
	return
}
func (ng *NegocioApp) GetStore() *storelayer.NegocioStore {
	return &ng.store
}

func  (ng *NegocioApp) decode(context *gin.Context){
	if err := context.BindJSON(ng.store.Single); err != nil {
		log.Println(err)
	}
    ng.store.Single.Init(0)
	return
}

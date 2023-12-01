package applayer

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/juanheza/facturador/helperlayer"
	"github.com/juanheza/facturador/modellayer"
	storelayer "github.com/juanheza/facturador/storeLayer"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NegocioApp struct {
	Response modellayer.Response
	store    storelayer.NegocioStore
}

var negocioApp NegocioApp

func (ng *NegocioApp) Create(context *gin.Context) {
	// Call BindJSON to bind the received JSON
	negocioApp.decode(context)
	userApp := &UserApp{}

	negocioApp.store.Single.GenerateBearer()
	user := userApp.CreateAdmin(context, negocioApp.store.Single.NegocioID)
	negocioApp.store.Single.Usuarios = append(negocioApp.store.Single.Usuarios, user)

	_, err := negocioApp.store.Create()
	if err != nil {
		return
	}
	negocioApp.Response = modellayer.Response{
		Id:      helperlayer.Success,
		Message: negocioApp.store.Single.ToJson(),
	}
	context.String(http.StatusOK, negocioApp.Response.ToJson())
}
func (ng *NegocioApp) Read(context *gin.Context) {
	id := context.Param("id")
	negocioApp.decode(context)
	if id != "" {
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			panic(err)
		}
		negocioApp.store.Single.NegocioID = objID
		negocioApp.store.Read()
	} else {
		negocioApp.store.ReadMany()
	}
	context.String(http.StatusOK, "Hello !!")
}
func (ng *NegocioApp) Update(context *gin.Context) {
	negocioApp.decode(context)
	negocioApp.store.Update()
	context.String(http.StatusOK, "Hello !!")
}
func (ng *NegocioApp) Delete(context *gin.Context) {
	negocioApp.decode(context)
	negocioApp.store.Delete()
	context.String(http.StatusOK, "Hello !!")
}
func (ng *NegocioApp) GetStore() *storelayer.NegocioStore {
	return &ng.store
}

func (ng *NegocioApp) decode(context *gin.Context) {
	negocio := modellayer.NewNegocio()
	if err := context.BindJSON(negocio); err != nil {
		log.Println(err)
	}
	if negocio.NegocioID == primitive.NilObjectID {
		negocio.NegocioID = primitive.NewObjectID()
	}
	negocio.Inicio = time.Now()
	negocio.Final = negocio.Inicio.AddDate(0, negocio.Periodo, 0)
	ng.store.Single = negocio
	ng.store.Single.Init(0)
	return
}

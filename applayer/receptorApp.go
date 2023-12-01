package applayer

import (
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/juanheza/facturador/modellayer"
	storelayer "github.com/juanheza/facturador/storeLayer"
)

type ReceptorApp struct {
	Response modellayer.Response
	store    storelayer.ReceptorStore
}

var receptorApp ReceptorApp

func (rc *ReceptorApp) Create(context *gin.Context) {
	receptorApp.decode(context)
	_, err := receptorApp.store.Create()
	if err != nil {
		return
	}
	return
}

func (rc *ReceptorApp) CreateAdmin(context *gin.Context, negocio primitive.ObjectID) {
	receptorApp.decode(context)
	_, err := receptorApp.store.Create()
	if err != nil {
		return
	}
	return
}

func (rc *ReceptorApp) Read(context *gin.Context) {
	receptorApp.decode(context)
	receptorApp.store.Read()
	return
}
func (rc *ReceptorApp) Update(context *gin.Context) {
	receptorApp.decode(context)
	receptorApp.store.Update()
	return
}
func (rc *ReceptorApp) Delete(context *gin.Context) {
	receptorApp.decode(context)
	receptorApp.store.Delete()
	return
}
func (rc *ReceptorApp) GetStore() *storelayer.ReceptorStore {
	return &rc.store
}

func (rc *ReceptorApp) decode(context *gin.Context) {
	receptor := &modellayer.Receptor{}
	if err := context.BindJSON(receptor); err != nil {
		log.Println(err)
	}
	rc.store.Single = receptor
	return
}

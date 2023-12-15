package applayer

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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
}

func (rc *ReceptorApp) CreateAdmin(context *gin.Context, negocio primitive.ObjectID) {
	receptorApp.decode(context)
	_, err := receptorApp.store.Create()
	if err != nil {
		return
	}
}

func (rc *ReceptorApp) Read(context *gin.Context) {
	receptorApp.decode(context)
	receptorApp.store.Read()
}
func (rc *ReceptorApp) Update(context *gin.Context) {
	receptorApp.decode(context)
	receptorApp.store.Update()
}
func (rc *ReceptorApp) Delete(context *gin.Context) {
	receptorApp.decode(context)
	receptorApp.store.Delete()
}
func (rc *ReceptorApp) GetStore() *storelayer.ReceptorStore {
	return &rc.store
}

func (rc *ReceptorApp) decode(context *gin.Context) {
	receptor := &modellayer.Receptor{}
	if err := context.ShouldBindBodyWith(receptor, binding.JSON); err != nil {
		log.Println(err)
	}
	rc.store.Single = receptor
}

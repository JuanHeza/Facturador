package applayer

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/juanheza/facturador/modellayer"
	storelayer "github.com/juanheza/facturador/storeLayer"
)

type NegocioApp struct {
	Response modellayer.Response
	store    storelayer.NegocioStore
}

func (ng *NegocioApp) Create(context *gin.Context) {
	// Call BindJSON to bind the received JSON
	negocioApp := decode(context)
	negocioApp.store.Create()
	return
}
func (ng *NegocioApp) Read(context *gin.Context) {
	negocioApp := decode(context)
	negocioApp.store.Read()
	return
}
func (ng *NegocioApp) Update(context *gin.Context) {
	negocioApp := decode(context)
	negocioApp.store.Update()
	return
}
func (ng *NegocioApp) Delete(context *gin.Context) {
	negocioApp := decode(context)
	negocioApp.store.Delete()
	return
}
func (ng *NegocioApp) GetStore() *storelayer.NegocioStore {
	return &ng.store
}

func decode(context *gin.Context) (ng NegocioApp) {
	if err := context.BindJSON(&ng.store.Single); err != nil {
		log.Println(err)
	}
	return
}

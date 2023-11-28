package applayer

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/juanheza/facturador/modellayer"
)

 type NegocioApp struct {
     Response modellayer.Response
     Data modellayer.Negocio `json:"-"`
 }

func (ng *NegocioApp)Create(context *gin.Context){
    // Call BindJSON to bind the received JSON
    negocioApp := decode(context)
    return
}
func (ng *NegocioApp)Read(context *gin.Context){
    negocioApp := decode(context)
    return
}
func (ng *NegocioApp)Update(context *gin.Context){
    negocioApp := decode(context)
    return
}
func (ng *NegocioApp)Delete(context *gin.Context){
    negocioApp := decode(context)
    return
}

func decode(context *gin.Context)(ng NegocioApp){
    if err := context.BindJSON(&ng.Data); err != nil {
        log.Println(err)
    }
    return
}
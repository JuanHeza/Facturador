package httplayer

import (
	"fmt"
	"net/http"
	_ "time"

	"github.com/gin-gonic/gin"
	"github.com/juanheza/facturador/applayer"
)

var negocioApp applayer.NegocioApp

func initNegocioRoutes(rt *gin.Engine) *gin.Engine {

	negocioApp := &applayer.NegocioApp{}
	/*
	   dummy := &modellayer.Negocio{NegocioID: 12, Clave: "evilPanda"}
	   dummy.GenerateBearer(time.Time{}, 1)
	   fmt.Println(dummy)
	*/
	rt.GET("/negocio/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK, fmt.Sprintf("Hello %s!!", name))
	})

	/*
	   PARAMS
	       Negocio
	           Clave
	           PeriodoVigencia
	           DiasVigencia
	           Folios
	           RazonSocial
	           Rfc
	           RegimenFiscal
	           CodigoPostal
	       User
	           Correo
	*/
	rt.POST("/negocio", negocioApp.Create)

	rt.PUT("/negocio", negocioApp.Update)

	rt.GET("/negocio/:id", negocioApp.Read)

	rt.DELETE("/negocio", negocioApp.Delete)

	return rt
}

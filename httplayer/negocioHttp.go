package httplayer

import (
	
	_ "time"

	"github.com/gin-gonic/gin"
	"github.com/juanheza/facturador/applayer"
)

var negocioApp applayer.NegocioApp

func initNegocioRoutes(rt *gin.Engine) *gin.Engine {

	negocioApp := applayer.NewNegocioApp()

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
	rt.POST("/negocio", negocioApp.Create) // privacidad maxima
    
	rt.PUT("/negocio", negocioApp.Update) // administradores

	rt.GET("/negocio", negocioApp.Read) //publico con params para data

	rt.DELETE("/negocio", negocioApp.Delete) 

	rt.PUT("/negocio/:id", negocioApp.Update) // administradores

	rt.GET("/negocio/:id", negocioApp.Read) //publico con params para data

	rt.DELETE("/negocio/:id", negocioApp.Delete) // privacidad maxima

	return rt
}

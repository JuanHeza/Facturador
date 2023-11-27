package httplayer

import (
	"fmt"
	"net/http"
	
	"github.com/gin-gonic/gin"
	_"github.com/juanheza/facturador/modellayer"
)

func initNegocioRoutes(rt *gin.Engine) *gin.Engine {

	rt.GET("/negocio/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK, fmt.Sprintf("Hello %s!!", name))
	})

	rt.POST("/negocio/Crear", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello !!")
	})

	rt.PUT("/negocio/Editar", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello !!")
	})

	rt.GET("/negocio/Leer", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello !!")
	})

	rt.DELETE("/negocio/Eliminar", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello !!")
	})

	return rt
}

package httplayer

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/juanheza/facturador/modellayer"
)

func initReceptorRoutes(rt *gin.Engine) *gin.Engine {

	rt.GET("/receptor/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK, fmt.Sprintf("Hello %s!!", name))
	})

	rt.POST("/receptor/Crear", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello !!")
	})

	rt.PUT("/receptor/Editar", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello !!")
	})

	rt.GET("/receptor/Leer", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello !!")
	})

	rt.DELETE("/receptor/Eliminar", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello !!")
	})

	return rt
}

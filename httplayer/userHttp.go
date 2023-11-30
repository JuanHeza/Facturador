package httplayer

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/juanheza/facturador/modellayer"
)

func initUserRoutes(rt *gin.Engine) *gin.Engine {

	rt.GET("/User/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK, fmt.Sprintf("Hello %s!!", name))
	})

	rt.POST("/User/Crear", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello !!")
	})

	rt.PUT("/User/Editar", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello !!")
	})

	rt.GET("/User/Leer", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello !!")
	})

	rt.DELETE("/User/Eliminar", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello !!")
	})

	return rt
}

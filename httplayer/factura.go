package httplayer

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func initFacturasRoutes(rt *gin.Engine) *gin.Engine {
	
	rt.GET("/factura/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK, fmt.Sprintf("Hello %s!!", name))
	})

	rt.POST("/factura/Crear", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello !!")
	})

	rt.GET("/factura/Leer", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello !!")
	})

	rt.DELETE("/factura/Eliminar", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello !!")
	})
	return rt
}

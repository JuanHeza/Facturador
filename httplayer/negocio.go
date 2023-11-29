package httplayer

import (
	"fmt"
	"net/http"
	_ "time"

	"github.com/gin-gonic/gin"
	"github.com/juanheza/facturador/applayer"
)

func initNegocioRoutes(rt *gin.Engine) *gin.Engine {
	/*
	   dummy := &modellayer.Negocio{NegocioID: 12, Clave: "evilPanda"}
	   dummy.GenerateBearer(time.Time{}, 1)
	   fmt.Println(dummy)
	*/
	rt.GET("/negocio/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK, fmt.Sprintf("Hello %s!!", name))
	})

	rt.POST("/negocio", func(context *gin.Context) {
		negocioApp := &applayer.NegocioApp{}
		negocioApp.Create(context)
		context.String(http.StatusOK, "OK")
	})

	rt.PUT("/negocio", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello !!")
	})

	rt.GET("/negocio", func(context *gin.Context) {
		negocioApp := &applayer.NegocioApp{}
		negocioApp.Read(context)
		context.String(http.StatusOK, "Hello !!")
	})

	rt.DELETE("/negocio", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello !!")
	})

	return rt
}

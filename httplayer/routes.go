package httplayer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juanheza/facturador/helperlayer"
	"github.com/juanheza/facturador/modellayer"
)

func New() {
	router := gin.Default()
	router.Use(Logger())

	router.GET("/user/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK, fmt.Sprintf("Hello %s!!", name))
	})

	router.GET("/auth/badRequest", func(context *gin.Context) {

		context.String(http.StatusOK, "Bad Request")
	})

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello world!")
	})

	initFacturasRoutes(router)

	initNegocioRoutes(router)

	initReceptorRoutes(router)

	initUserRoutes(router)

	// starts the server at port 8080
	router.Run(":8080")
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.FullPath() != "/auth/badRequest" {
			t := time.Now()

			// Set example variable
			c.Set("example", "12345")

			// before request

			if checkAuth(c.Request.Header["Authorization"]) {
				c.Next()
			} else {
				c.Redirect(http.StatusMovedPermanently, "/auth/badRequest")
				c.Abort()
			}
			// after request
			latency := time.Since(t)
			log.Print(latency)

			// access the status we are sending
			status := c.Writer.Status()
			log.Println(status)
		}
	}
}

func checkAuth(auth []string) bool {
	fmt.Println(auth)
	if len(auth) != 1 {
		return false
	}
	token := strings.Trim(auth[0], "Bearer ")
	jsonStr := helperlayer.Decrypt(token)
	data := &modellayer.Token{}

	json.Unmarshal([]byte(jsonStr), &data)

	return data.Validate()
}

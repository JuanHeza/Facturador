package applayer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
	"github.com/juanheza/facturador/helperlayer"
	"github.com/juanheza/facturador/modellayer"
	_ "go.mongodb.org/mongo-driver/bson/primitive"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func GetTestGinContext() *gin.Context {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}
	return ctx
}

func buildBody(i interface{}, c *gin.Context) *gin.Context {
	body, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	// the request body must be an io.ReadCloser
	// the bytes buffer though doesn't implement io.Closer,
	// so you wrap it in a no-op closer
	c.Request, _ = http.NewRequest("Get", "/negocio", bytes.NewBuffer(body))
	return c
}

func TestNegocioApp_Create(t *testing.T) {
	type args struct {
		context *gin.Context
		data    map[string]interface{}
	}
	tests := []struct {
		name string
		ng   *NegocioApp
		args args
	}{
		{
			name: "Prueba 1",
			ng:   &NegocioApp{},
			args: args{
				context: GetTestGinContext(),
				data: map[string]interface{}{
					"correo":          "juan@heza.com",
					"clave":           "billingbull",
					"colorPrincipal":  "#defeca",
					"colorSecundario": "#cebada",
					"configuracionNegocio": map[string]interface{}{
						"periodoVigencia": helperlayer.Dias,
						"diasVigencia":    10,
						"folio":           20,
						"vigencia": map[string]interface{}{
							"periodo": 12,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		tt.args.context = buildBody(tt.args.data, tt.args.context)

		t.Run(tt.name, func(t *testing.T) {
			tt.ng.Create(tt.args.context)
		})
	}
}

func TestNegocioApp_Read(t *testing.T) {
	type args struct {
		result modellayer.Response
		data   string
	}
	tests := []struct {
		name string
		ng   *NegocioApp
		args args
	}{
		{
			name: "Prueba 1",
			ng:   &NegocioApp{},
			args: args{
				result: modellayer.Response{
					Id:      helperlayer.Success,
					Message: "Negocios",
				},
				data: "/negocio/",
			},
		},
		{
			name: "Prueba 2",
			ng:   &NegocioApp{},
			args: args{
				result: modellayer.Response{
					Id:      helperlayer.Success,
					Message: "Negocios",
				},
				data: "/negocio/656980b5b01c2801819367da",
			},
		},
	}

	r := SetUpRouter()
	negocioApp := NewNegocioApp()
	r.GET("/negocio/", negocioApp.Read)
	r.GET("/negocio/:id", negocioApp.Read)

	for _, tt := range tests {
		fmt.Println(tt.name)
		req, _ := http.NewRequest("GET", tt.args.data, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		res := modellayer.Response{}
		json.NewDecoder(w.Body).Decode(&res)
		assert.Equal(t, tt.args.result.Id, res.Id)
		assert.Equal(t, tt.args.result.Message, res.Message)
		fmt.Println(res.Data)
		assert.Equal(t, http.StatusOK, w.Code)
	}
}

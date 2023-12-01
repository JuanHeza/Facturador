package applayer

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/juanheza/facturador/helperlayer"
	"github.com/juanheza/facturador/modellayer"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
	jsonStr, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}

	// the request body must be an io.ReadCloser
	// the bytes buffer though doesn't implement io.Closer,
	// so you wrap it in a no-op closer
	c.Request, _ = http.NewRequest("Get", "/negocio", bytes.NewBuffer(jsonStr))
	return c
}

func TestNegocioApp_Create(t *testing.T) {
	type args struct {
		context *gin.Context
		data    modellayer.Negocio
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
				data: modellayer.Negocio{
					NegocioID:       primitive.NewObjectID(),
					Clave:           "billingbull",
					ColorPrincipal:  "#defeca",
					ColorSecundario: "#cebada",
					ConfiguracionNegocio: modellayer.ConfiguracionNegocio{
						PeriodoVigencia: helperlayer.Dias,
						DiasVigencia:    10,
						Folios:          20,
						Vigencia: modellayer.Vigencia{
							Periodo: 12,
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

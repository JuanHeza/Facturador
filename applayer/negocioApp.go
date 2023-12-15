package applayer

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/juanheza/facturador/helperlayer"
	"github.com/juanheza/facturador/modellayer"
	storelayer "github.com/juanheza/facturador/storeLayer"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NegocioApp struct {
	Response modellayer.Response
	store    storelayer.NegocioStore
}

func NewNegocioApp() *NegocioApp {
	return &NegocioApp{
		Response: modellayer.NewResponse(),
		store:    storelayer.NewNegocioStore(),
	}
}

func (ng *NegocioApp) Clean() gin.HandlerFunc {
	return func(c *gin.Context) {
		ng.store = storelayer.NewNegocioStore()
	}
}

func (ng *NegocioApp) Create(context *gin.Context) {
	// Call BindJSON to bind the received JSON

	ng.decode(context)
	userApp := NewUserApp()
	if err := ng.ValidateCreation(); err != nil{
		ng.Response = modellayer.Response{
			Id:      helperlayer.Error,
			Message: err.Error(),
		}
	}else{ 
		ng.store.Single.GenerateBearer()
		user := userApp.CreateAdmin(context, ng.store.Single.NegocioID)
		ng.store.Single.Owner = *user

		_, err := ng.store.Create()
		if err != nil {
			log.Println(err)
			return
		}
		ng.Response = modellayer.Response{
			Id:      helperlayer.Success,
			Message: "Negocio Creado",
		}
		ng.Response.Set("Single", ng.store.Single)
	}
	context.String(http.StatusOK, ng.Response.ToJson())
}

func (ng *NegocioApp) Read(context *gin.Context) {
	id := context.Param("id")
	var resultError error
	//ng.decode(context)
	if id != "" {
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			panic(err)
		}
		ng.store.Single.NegocioID = objID
		resultError = ng.store.Read()
	} else {
		resultError = ng.store.ReadMany()
	}
	ng.Response = modellayer.Response{
		Id:      helperlayer.Success,
		Message: "Negocios",
	}
	if resultError != nil {
		ng.Response.Id = helperlayer.Error
		ng.Response.Message = resultError.Error()
	}
	ng.Response.Set("List", ng.store.List)
	ng.Response.Set("Single", ng.store.Single)
	context.String(http.StatusOK, ng.Response.ToJson())
}

func (ng *NegocioApp) Update(context *gin.Context) {
	ng.decode(context)
	ng.store.Update()
	context.String(http.StatusOK, "Hello !!")
}

func (ng *NegocioApp) Delete(context *gin.Context) {
	ng.decode(context)
	ng.store.Delete()
	context.String(http.StatusOK, "Hello !!")
}

func (ng *NegocioApp) GetStore() *storelayer.NegocioStore {
	return &ng.store
}

func (ng *NegocioApp) decode(context *gin.Context) {
	negocio := modellayer.NewNegocio()
	if err := context.ShouldBindBodyWith(negocio, binding.JSON); err != nil {
		log.Println(err)
	}
	if negocio.NegocioID == primitive.NilObjectID {
		negocio.NegocioID = primitive.NewObjectID()
	}

	auxInicio := time.Now()
	negocio.Inicio = &auxInicio
	auxFinal := negocio.Inicio.AddDate(0, negocio.Periodo, 0)
	negocio.Final = &auxFinal
	ng.store.Single = negocio
	ng.store.Single.Init(0)
}

func (ng *NegocioApp) ValidateCreation() (err error){
	single := ng.store.Single
	if single.Clave == "" || single.Owner.Correo == "" || single.PeriodoVigencia == helperlayer.InvalidVigencia || single.Folios == 0 || single.Periodo == 0{
		err = errors.New("Bad request: Incomplete data")
		return
	}
	return
}
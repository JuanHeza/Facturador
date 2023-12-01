package applayer

import (
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/juanheza/facturador/modellayer"
	storelayer "github.com/juanheza/facturador/storeLayer"
)

type UserApp struct {
	Response modellayer.Response
	store    storelayer.UserStore
}

var userApp UserApp

func (us *UserApp) Create(context *gin.Context) {
	userApp.decode(context)
	_, err := userApp.store.Create()
	if err != nil {
		return
	}
	return
}

func (us *UserApp) CreateAdmin(context *gin.Context, negocio primitive.ObjectID) *modellayer.User {
	userApp.decode(context)
	userApp.store.Single.SetAdmin(negocio)
	_, err := userApp.store.Create()
	if err != nil {
		return nil
	}
	return userApp.store.Single
}

func (us *UserApp) Read(context *gin.Context) {
	userApp.decode(context)
	userApp.store.Read()
	return
}
func (us *UserApp) Update(context *gin.Context) {
	userApp.decode(context)
	userApp.store.Update()
	return
}
func (us *UserApp) Delete(context *gin.Context) {
	userApp.decode(context)
	userApp.store.Delete()
	return
}
func (us *UserApp) GetStore() *storelayer.UserStore {
	return &us.store
}

func (us *UserApp) decode(context *gin.Context) {
	user := modellayer.NewUser()
	if err := context.BindJSON(user); err != nil {
		log.Println(err)
	}
	if user.UsuarioID == primitive.NilObjectID {
		user.UsuarioID = primitive.NewObjectID()
	}
	us.store.Single = user
	return
}

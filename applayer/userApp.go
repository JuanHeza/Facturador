package applayer

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/juanheza/facturador/modellayer"
	storelayer "github.com/juanheza/facturador/storeLayer"
)

type UserApp struct {
	Response modellayer.Response
	store    storelayer.UserStore
}

func NewUserApp() *UserApp {
	return &UserApp{
		Response: modellayer.NewResponse(),
		store:    storelayer.NewUserStore(),
	}
}

func (us *UserApp) Create(context *gin.Context) {
	us.decode(context)
	_, err := us.store.Create()
	if err != nil {
		return
	}
}

func (us *UserApp) CreateAdmin(context *gin.Context, negocio primitive.ObjectID) *modellayer.User {
	us.decode(context)
	us.store.Single.SetAdmin(negocio)
	_, err := us.store.Create()
	if err != nil {
		return nil
	}
	return us.store.Single
}

func (us *UserApp) Read(context *gin.Context) {
	us.decode(context)
	us.store.Read()
}
func (us *UserApp) Update(context *gin.Context) {
	us.decode(context)
	us.store.Update()
}
func (us *UserApp) Delete(context *gin.Context) {
	us.decode(context)
	us.store.Delete()
}
func (us *UserApp) GetStore() *storelayer.UserStore {
	return &us.store
}

func (us *UserApp) decode(context *gin.Context) {
	user := modellayer.NewUser()
	if err := context.ShouldBindBodyWith(user, binding.JSON); err != nil {
		log.Println(err)
	}
	if user.UsuarioID == primitive.NilObjectID {
		user.UsuarioID = primitive.NewObjectID()
	}
	us.store.Single = user
}

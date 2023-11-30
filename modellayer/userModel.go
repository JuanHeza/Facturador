package modellayer

import (
	"github.com/juanheza/facturador/helperlayer"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct{
	UsuarioID primitive.ObjectID `bson:"_id,omitempty"`
	Usuario string `bson:"usuario,omitempty"`
	Password string `bson:"password,omitempty"`
	Nombre string `bson:"nombre,omitempty"`
	ApellidoPaterno string `bson:"apellidoPaterno,omitempty"`
	ApellidoMaterno string `bson:"apellidoMaterno,omitempty"`
	Correo string `bson:"correo,omitempty"`
	PerfilID helperlayer.Perfil `bson:"perfil,omitempty"`
    Sucursales []primitive.ObjectID `bson:"sucursales,omitempty"`
    Negocio primitive.ObjectID `bson:"negocio,omitempty"`
	Control
}

func (us *User) SetAdmin(negocio primitive.ObjectID){
    us.Usuario = "Administrador"
    us.Nombre = "Administrador"
    us.ApellidoPaterno = "Negocio"
    us.PerfilID = helperlayer.Admin
    us.Negocio = negocio
    us.Password = helperlayer.GeneratePassword()
    return
}
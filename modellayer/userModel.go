package modellayer

import (
	"encoding/json"
	"fmt"

	"github.com/juanheza/facturador/helperlayer"
	"go.mongodb.org/mongo-driver/bson"
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
	PerfilID helperlayer.Perfil `bson:"perfilId,omitempty"`
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


func (us *User) ToFilter() (filter bson.M){
    data, err := bson.Marshal(us)
    if err != nil {
        return
    }
    err = bson.Unmarshal(data, &filter)
    filter = bson.M{"_id":us.UsuarioID}
    return
}

func NewUser() *User {
    return &User{
        UsuarioID: primitive.NewObjectID(),
        Control: NewControl(),
    }
}

func (us *User) ToJson() (output string) {
    bits, err := json.Marshal(us)
    if err != nil {
        fmt.Printf("Error: %s", err)
        return
    }
    output = string(bits)
    return
}
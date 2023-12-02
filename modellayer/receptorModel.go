package modellayer

import (
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type Receptor struct {
	ReceptorID     primitive.ObjectID
	RazonSocial    string
	Rfc            string
	Correo         string
	RegimenFiscal  string
	UsoCFDI        string
	TipoVenta      string
	Calle          string
	NumeroExterior int
	NumeroInterior int
	Colonia        string
	CodigoPostal   int
	Municipio      string
	Estado         string
	Pais           string
	Control
}

func NewReceptor() *Receptor {
    return &Receptor{
		ReceptorID: primitive.NewObjectID(),
        Control: NewControl(),
    }
}

func (rc *Receptor) ToFilter() (filter bson.M){
    data, err := bson.Marshal(rc)
    if err != nil {
        return
    }
    err = bson.Unmarshal(data, &filter)
    filter = bson.M{"_id":rc.ReceptorID}
    return
}

func (rc *Receptor) ToJson() (output string) {
    bits, err := json.Marshal(rc)
    if err != nil {
        fmt.Printf("Error: %s", err)
        return
    }
    output = string(bits)
    return
}
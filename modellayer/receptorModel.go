package modellayer
import (
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
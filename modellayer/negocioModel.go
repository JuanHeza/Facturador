package modellayer

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/juanheza/facturador/helperlayer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Negocio struct {
	NegocioID            primitive.ObjectID `json:",omitempty" bson:"_id,omitempty"`
	BearerToken          string             `json:"-,omitempty" bson:"bearer_token,omitempty"`
	Clave                string             `json:",omitempty" bson:"clave,omitempty"`
	Logo                 string             `json:",omitempty" bson:"logo,omitempty"`
	ColorPrincipal       string             `json:",omitempty" bson:"color_primario,omitempty"`
	ColorSecundario      string             `json:",omitempty" bson:"color_secundario,omitempty"`
	ConfiguracionNegocio `json:",omitempty"`
	DatosGenerales       `bson:",omitempty" json:",omitempty"`
	Emisor               `bson:",omitempty" json:",omitempty"`
	Sucursales           []*Sucursal `json:",omitempty" bson:"sucursales,omitempty"`
	Owner                User        `json:",omitempty" bson:"owner,omitempty"`
	Control              `json:",omitempty"`
}

type ConfiguracionNegocio struct {
	PeriodoVigencia helperlayer.Vigencia `json:",omitempty"  bson:"periodo_vigencia,omitempty"`
	DiasVigencia    int                  `json:",omitempty"  bson:"dias_vigencia,omitempty"`
	Folios          int                  `json:",omitempty"  bson:"folios,omitempty"`
	Vigencia        `bson:",omitempty"`
}

type Vigencia struct {
	Periodo int        `json:",omitempty" bson:",omitempty"`
	Inicio  *time.Time `json:",omitempty" bson:"inicio,omitempty"`
	Final   *time.Time `json:",omitempty" bson:"final,omitempty"`
}

type Link struct {
	Nombre string `json:",omitempty"  bson:"nombre,omitempty"`
	Url    string `json:",omitempty"  bson:"url,omitempty"`
}

type Post struct {
	Imagen string `json:",omitempty"  bson:"imagen,omitempty"`
	Link
}

type Emisor struct {
	RazonSocial   string `json:",omitempty"  bson:"razon_social,omitempty"`
	Rfc           string `json:",omitempty"  bson:"rfc,omitempty"`
	RegimenFiscal string `json:",omitempty"  bson:"regimen_fiscal,omitempty"`
	CodigoPostal  int    `json:",omitempty"  bson:"codigo_postal,omitempty"`
	Csd           string `json:",omitempty"  bson:"csd,omitempty"`
	Password      string `json:",omitempty"  bson:"password,omitempty"`
	Key           string `json:",omitempty"  bson:"key,omitempty"`
}

func (df *Emisor) SetNull() {
	df = &Emisor{}
}

type DatosGenerales struct {
	Titulo            string   `json:",omitempty"  bson:"titulo,omitempty"`
	ImagenCabecera    string   `json:",omitempty"  bson:"imagen_cabecera,omitempty"`
	Menu              []Link   `json:",omitempty"  bson:"menu,omitempty"`
	DescripcionTitulo string   `json:",omitempty"  bson:"descripcion_titulo,omitempty"`
	DescripcionTexto  string   `json:",omitempty"  bson:"descripcion_texto,omitempty"`
	Pie               []Post   `json:",omitempty"  bson:"pie,omitempty"`
	Direccion         string   `json:",omitempty"  bson:"direccion,omitempty"`
	Telefono          []string `json:",omitempty"  bson:"telefono,omitempty"`
}

func (dg *DatosGenerales) SetNull() {
	dg = &DatosGenerales{}
}

func (ng *Negocio) GenerateBearer() {
	auxInicio := time.Now()
	ng.Inicio = &auxInicio
	auxFinal := ng.Inicio.AddDate(0, ng.Periodo, 0)
	ng.Final = &auxFinal

	bearerData := &Token{
		NegocioId: ng.NegocioID.Hex(),
		Clave:     ng.Clave,
		Inicio:    ng.Inicio.Unix(),
		Fin:       ng.Final.Unix(),
	}
	plaintext, err := json.Marshal(bearerData)
	if err != nil {
		return
	}
	ng.BearerToken = helperlayer.Encrypt(plaintext)
}

func (ng *Negocio) ToJson() (output string) {
	bits, err := json.Marshal(ng)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	output = string(bits)
	return
}

func NewNegocio() *Negocio {
	return &Negocio{
		NegocioID: primitive.NewObjectID(),
		Control:   NewControl(),
	}
}

func (ng *Negocio) ToFilter() (filter bson.M) {
	data, err := bson.Marshal(ng)
	if err != nil {
		return
	}
	err = bson.Unmarshal(data, &filter)
	filter = bson.M{"_id": ng.NegocioID}
	return
}

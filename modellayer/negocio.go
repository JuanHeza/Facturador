package modellayer

// https://blog.reedsy.com/fantasy-map-generators/
import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/juanheza/facturador/helperlayer"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Negocio struct {
	NegocioID       primitive.ObjectID   `json:",omitempty" bson:"_id,omitempty"`
	BearerToken     string               `json:"-" bson:"bearer_token,omitempty"`
	Clave           string               `json:",omitempty"  bson:"clave,omitempty"`
    Logo            string               `json:",omitempty"  bson:"logo,omitempty"`
	ColorPrincipal  string               `json:",omitempty"  bson:"color_primario,omitempty"`
	ColorSecundario string               `json:",omitempty"  bson:"color_secundario,omitempty"`
	PeriodoVigencia helperlayer.Vigencia `json:",omitempty"  bson:"periodo_vigencia,omitempty"`
	DiasVigencia    int                  `json:",omitempty"  bson:"dias_vigencia,omitempty"`
    Folios          int                  `json:",omitempty"  bson:"folios,omitempty"`
	DatosGenerales                       `bson:",omitempty"`
	DatosFiscales                        `bson:",omitempty"`
	Control
}

type Link struct {
	Nombre string `json:",omitempty"  bson:"nombre,omitempty"`
	Url    string `json:",omitempty"  bson:"url,omitempty"`
}

type Post struct {
	Imagen string `json:",omitempty"  bson:"imagen,omitempty"`
	Link
}

type DatosFiscales struct {
	RazonSocial   string `json:",omitempty"  bson:"razon_social,omitempty"`
	Rfc           string `json:",omitempty"  bson:"rfc,omitempty"`
	RegimenFiscal string `json:",omitempty"  bson:"regimen_fiscal,omitempty"`
	CodigoPostal  int    `json:",omitempty"  bson:"codigo_postal,omitempty"`
}

func (df *DatosFiscales) SetNull() {
	df = &DatosFiscales{}
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

func (ng *Negocio) GenerateBearer(tm time.Time, period int) {
	if tm.IsZero() {
		tm = time.Now()
	}
	fmt.Println(tm)
	bearerData := &Token{
		NegocioId: ng.NegocioID.Hex(),
		Clave:     ng.Clave,
		Inicio:    tm.Unix(),
		Fin:       tm.AddDate(0, period, 0).Unix(),
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

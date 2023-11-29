package modellayer
// https://blog.reedsy.com/fantasy-map-generators/
import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/juanheza/facturador/helperlayer"
)

type Negocio struct {
	NegocioID         int  `json:",omitempty"`
	BearerToken       string `json:"-"`
	Clave             string `json:",omitempty"`
	ColorPrincipal    string `json:",omitempty"`
	ColorSecundario   string `json:",omitempty"`
	PeriodoVigencia   helperlayer.Vigencia `json:",omitempty"`
	DiasVigencia      int `json:",omitempty"`
    DatosGenerales
	DatosFiscales
	Control
}

type Link struct {
	Nombre string `json:",omitempty"`
    Url string `json:",omitempty"`
}

type Post struct {
	Imagen string `json:",omitempty"`
	Link
}

type DatosFiscales struct {
	RazonSocial   string `json:",omitempty"`
	Rfc           string `json:",omitempty"`
	RegimenFiscal string `json:",omitempty"`
	CodigoPostal  int `json:",omitempty"`
}

func (df *DatosFiscales) SetNull(){
    df = &DatosFiscales{}
}

type DatosGenerales struct{
	Titulo            string `json:",omitempty"`
	Logo              string `json:",omitempty"`
	ImagenCabecera    string `json:",omitempty"`
	Menu              []Link `json:",omitempty"`
	DescripcionTitulo string `json:",omitempty"`
	DescripcionTexto  string `json:",omitempty"`
	Pie               []Post `json:",omitempty"`
	Direccion         string `json:",omitempty"`
	Telefono          []string `json:",omitempty"`    
}

func (dg *DatosGenerales) SetNull(){
    dg = &DatosGenerales{}
}

func (ng *Negocio) GenerateBearer(tm time.Time, period int) {
	if tm.IsZero() {
		tm = time.Now()
	}
	fmt.Println(tm)
	bearerData := &Token{
		NegocioId: ng.NegocioID,
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

func (ng *Negocio) ToJson() (output string){
    bits, err := json.Marshal(ng)
    if err != nil {
        fmt.Printf("Error: %s", err)
        return;
    }
    output = string(bits)
    return
}
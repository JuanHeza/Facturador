package modellayer

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/juanheza/facturador/helperlayer"
)

type Negocio struct {
	NegocioID         int
	BearerToken       string
	Titulo            string
	Clave             string
	Logo              string
	ImagenCabecera    string
	Menu              []Link
	ColorPrincipal    string
	ColorSecundario   string
	DescripcionTitulo string
	DescripcionTexto  string
	Pie               []Post
	Direccion         string
	Telefono          []string
	PeriodoVigencia   helperlayer.Vigencia
	DiasVigencia      int
	DatosFiscales
	Control
}

type Link struct {
	Nombre, Url string
}
type Post struct {
	Imagen string
	Link
}

type DatosFiscales struct {
	RazonSocial   string
	Rfc           string
	RegimenFiscal string
	CodigoPostal  int
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

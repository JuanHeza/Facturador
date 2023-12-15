package modellayer

import (
	"encoding/json"
	"time"

	"github.com/juanheza/facturador/helperlayer"
)

type Token struct {
	NegocioId string `json:",omitempty"`
	Clave     string `json:",omitempty"`
	Inicio    int64  `json:",omitempty"`
	Fin       int64  `json:",omitempty"`
}

func (tk *Token) Validate() bool {
	today := time.Now()
	inicio := time.Unix(tk.Inicio, 0)
	fin := time.Unix(tk.Fin, 0)
	return inicio.Before(today) && fin.After(today)
}

func NewToken(ng *Negocio) (bearerData *Token) {
	auxInicio := time.Now()
	ng.Inicio = &auxInicio
	auxFinal := ng.Inicio.AddDate(0, ng.Periodo, 0)
	ng.Final = &auxFinal

	bearerData = &Token{
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
	return
}

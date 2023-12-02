package modellayer

import (
	"time"
    "encoding/json"
    "github.com/juanheza/facturador/helperlayer"
)

type Token struct {
	NegocioId string
	Clave     string
	Inicio    int64
	Fin       int64
}

func (tk *Token) Validate() bool {
	today := time.Now()
	inicio := time.Unix(tk.Inicio, 0)
	fin := time.Unix(tk.Fin, 0)
	return inicio.Before(today) && fin.After(today)
}

func NewToken(ng *Negocio) (bearerData *Token) {
    ng.Inicio = time.Now()
    ng.Final = ng.Inicio.AddDate(0, ng.Periodo, 0)

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
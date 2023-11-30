package modellayer

import (
	"log"
	"time"
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
	log.Println(inicio)
	log.Println(today)
	log.Println(fin)
	return inicio.Before(today) && fin.After(today)
}

package modellayer

import (
	"time"

	"github.com/juanheza/facturador/helperlayer"
)

type Control struct {
	Estatus   helperlayer.Estatus `json:"-"`
	Eliminado helperlayer.Estatus `json:"-"`
	Creacion  ControlData `json:"-"`
	Edicion   ControlData `json:"-"`
}

type ControlData struct {
	Usuario int
	Fecha   time.Time
}

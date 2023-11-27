package modellayer

import (
	"time"

	"github.com/juanheza/facturador/helperlayer"
)

type Control struct {
	Estatus   helperlayer.Estatus
	Eliminado helperlayer.Estatus
	Creacion  ControlData
	Edicion   ControlData
}

type ControlData struct {
	Usuario int
	Fecha   time.Time
}

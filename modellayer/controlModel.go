package modellayer

import (
	"time"

	"github.com/juanheza/facturador/helperlayer"
)

type Control struct {
	Estatus   helperlayer.Estatus `json:",omitempty"  bson:",omitempty"`
	Eliminado helperlayer.Estatus `json:"-"  bson:",omitempty"`
	Creacion  ControlData         `json:",omitempty"  bson:",omitempty"`
	Edicion   ControlData         `json:",omitempty"  bson:",omitempty"`
}

type ControlData struct {
	Usuario int       `json:",omitempty"`
	Fecha   *time.Time `json:",omitempty"`
}

func (cd *ControlData) update(user int) {
	cd.Usuario = user
	aux := time.Now()
	cd.Fecha = &aux
}

func (c *Control) SetEstatus(estatus helperlayer.Estatus, user int) {
	c.Estatus = estatus
	c.Edicion.update(user)
}
func (c *Control) SetEliminado(eliminado helperlayer.Estatus, user int) {
	c.Eliminado = eliminado
	c.Edicion.update(user)
}

func (c *Control) Init(user int) {
	c.Estatus = helperlayer.Activo
	c.Eliminado = helperlayer.NoEliminado
	c.Edicion.update(user)
	c.Creacion.update(user)
}

func NewControl() Control {
	control := Control{}
	control.Init(0)
	return control
}

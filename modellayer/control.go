package modellayer

import (
	"time"

	"github.com/juanheza/facturador/helperlayer"
)

type Control struct {
	Estatus   helperlayer.Estatus `json:",omitempty"  bson:",omitempty"`
	Eliminado helperlayer.Estatus `json:"-,omitempty"  bson:",omitempty"`
	Creacion  ControlData `json:",omitempty"  bson:",omitempty"`
	Edicion   ControlData `json:",omitempty"  bson:",omitempty"`
}

type ControlData struct {
	Usuario int
	Fecha   time.Time
}

func( cd *ControlData) update(user int){
    cd.Usuario = user
    cd.Fecha = time.Now()
}

func (c *Control) SetEstatus(estatus helperlayer.Estatus, user int){
    c.Estatus = estatus
    c.Edicion.update(user)
}
func (c *Control) SetEliminado(eliminado helperlayer.Estatus, user int){
    c.Eliminado = eliminado
    c.Edicion.update(user)
}

func (c *Control) Init(user int){
    c.Estatus = helperlayer.Activo
    c.Eliminado = helperlayer.NoEliminado
    c.Edicion.update(user)
    c.Creacion.update(user)
}
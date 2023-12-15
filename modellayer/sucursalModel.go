package modellayer

type Sucursal struct {
	SucursalId string  `json:",omitempty"`
	Clave      string  `json:",omitempty"`
	Nombre     string  `json:",omitempty"`
	Telefono   string  `json:",omitempty"`
	Tickets    []Venta `json:",omitempty"`
	Control    `json:",omitempty"`
}

func NewSucursal() *Sucursal {
	return &Sucursal{
		Control: NewControl(),
	}
}

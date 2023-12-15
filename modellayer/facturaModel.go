package modellayer

type Factura struct {
	FacturaId   int    `json:",omitempty"`
	RutaArcivo  string `json:",omitempty"`
	FacturaUUID string `json:",omitempty"`
	Emisor      `json:",omitempty"`
	Receptor    `json:",omitempty"`
	Tickets     []Venta `json:",omitempty"`
	Control     `json:",omitempty"`
}

func NewFactura() *Factura {
	return &Factura{
		Control: NewControl(),
	}
}

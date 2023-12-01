package modellayer

type Factura struct {
	FacturaId   int
	RutaArcivo  string
	FacturaUUID string
	Emisor
	Receptor
	Tickets []Venta
	Control
}

func NewFactura() *Factura {
    return &Factura{
        Control: NewControl(),
    }
}
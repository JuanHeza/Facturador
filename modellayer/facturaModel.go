package modellayer

type Factura struct {
	FacturaId   int
	RutaArcivo  string
	FacturaUUID string
	Negocio
	Receptor
	Venta
	Control
}

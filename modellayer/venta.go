package modellayer

import "time"

type Venta struct {
	VentaID    int
	NegocioID  int
	SucursalID int
	Folio      string
	Fecha      time.Time
	Conceptos []Item
	Total float32
	Control
}

type Item struct{
	ProductCode string
	Description string
	Unit string
	UnitCode string
	UnitPrice float32
	Quantity float32
	Subtotal float32
	Discount float32
    TaxObject string
    Taxes []Tax
}

type Tax struct{
	Total float32
    Name string
	Base float32
	Rate float32
    IsRetention bool
    IsQuota bool
}
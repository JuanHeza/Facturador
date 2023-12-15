package modellayer

import "time"

type Venta struct {
	VentaID    int       `json:"venta_id,omitempty"`
	NegocioID  int       `json:"negocio_id,omitempty"`
	SucursalID int       `json:"sucursal_id,omitempty"`
	Folio      string    `json:"folio,omitempty"`
	Fecha      time.Time `json:"fecha,omitempty"`
	Conceptos  []Item    `json:"conceptos,omitempty"`
	Total      float32   `json:"total,omitempty"`
	Control    `json:"control,omitempty"`
}

type Item struct {
	ProductCode string  `json:"product_code,omitempty"`
	Description string  `json:"description,omitempty"`
	Unit        string  `json:"unit,omitempty"`
	UnitCode    string  `json:"unit_code,omitempty"`
	UnitPrice   float32 `json:"unit_price,omitempty"`
	Quantity    float32 `json:"quantity,omitempty"`
	Subtotal    float32 `json:"subtotal,omitempty"`
	Discount    float32 `json:"discount,omitempty"`
	TaxObject   string  `json:"tax_object,omitempty"`
	Taxes       []Tax   `json:"taxes,omitempty"`
}

type Tax struct {
	Total       float32 `json:"total,omitempty"`
	Name        string  `json:"name,omitempty"`
	Base        float32 `json:"base,omitempty"`
	Rate        float32 `json:"rate,omitempty"`
	IsRetention bool    `json:"is_retention,omitempty"`
	IsQuota     bool    `json:"is_quota,omitempty"`
}


func NewVenta() *Venta {
    return &Venta{
        Control: NewControl(),
    }
}
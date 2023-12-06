package modellayer

type Sucursal struct{
    SucursalId string
    Clave string
    Nombre string
    Telefono string
    Tickets []Venta
    Control
}

func NewSucursal() *Sucursal {
    return &Sucursal{
        Control: NewControl(),
    }
}
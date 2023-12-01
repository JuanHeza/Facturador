package modellayer

type Sucursal struct{
    SucursalId string
    Clave string
    Nombre string
    Telefono string
    Tickets []Venta
    Control
}

// https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/read-operations/project/

func NewSucursal() *Sucursal {
    return &Sucursal{
        Control: NewControl(),
    }
}
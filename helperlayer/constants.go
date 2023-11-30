package helperlayer

import "fmt"

type Estatus int
type Vigencia int
type Perfil int
type Response string
type Collections string

const (
	Activo Estatus = iota + 1
	Inactivo
	Eliminado
	NoEliminado
    Facturado
    NoFacturado
    Pendiente

	Dias Vigencia = iota + 1
	Mes

	Admin Perfil = iota + 1

    Success Response = "SUCCESS"
    Error Response = "ERROR"

    PasswordSize = 8

    User Collections = "user"
    Negocio Collections = "negocio"
    Receptor Collections = "receptor"
)

func (cl Collections) ToString() string{
    return fmt.Sprint(cl)
}
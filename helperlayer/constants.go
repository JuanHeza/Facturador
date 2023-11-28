package helperlayer

type Estatus int
type Vigencia int
type Perfil int
type Response string
const (
	Activo Estatus = iota + 1
	Inactivo
	Eliminado
	NoEliminado

	Dias Vigencia = iota + 1
	Mes

	Admin Perfil = iota + 1

    Success Response = "SUCCESS"
    Error Response = "ERROR"

)
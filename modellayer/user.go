package modellayer

import "github.com/juanheza/facturador/helperlayer"

type User struct{
	UsuarioID int
	Usuario string
	Password string
	Nombre string
	ApellidoPaterno string
	ApellidoMaterno string
	Correo string
	PerfilID helperlayer.Perfil
	Control
}
package modellayer

type Receptor struct {
	ReceptorID     int
	RazonSocial    string
	Rfc            string
	Correo         string
	RegimenFiscal  string
	UsoCFDI        string
	TipoVenta      string
	Calle          string
	NumeroExterior int
	NumeroInterior int
	Colonia        string
	CodigoPostal   int
	Municipio      string
	Estado         string
	Pais           string
	Control
}

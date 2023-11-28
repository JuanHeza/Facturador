package modellayer

import "github.com/juanheza/facturador/helperlayer"


type Response struct{
    Id helperlayer.Response
    Message string
    Data map[string]any
}

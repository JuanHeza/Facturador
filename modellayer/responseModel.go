package modellayer

import (
	"encoding/json"
	"fmt"

	"github.com/juanheza/facturador/helperlayer"
)

type Response struct {
	Id      helperlayer.Response
	Message string
	Data    map[string]any
}
func NewResponse() Response{
    return Response{
        Data: map[string]any{},
    }
}

func (ng *Response) Set(key string, value any) {
    if ng.Data == nil {
        ng.Data = make(map[string]any)
    }
    ng.Data[key] = value
}

func (ng *Response) ToJson() (output string) {
	bits, err := json.Marshal(ng)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	output = string(bits)
	return
}

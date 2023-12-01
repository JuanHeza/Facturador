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

func (ng *Response) ToJson() (output string) {
	bits, err := json.Marshal(ng)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	output = string(bits)
	return
}

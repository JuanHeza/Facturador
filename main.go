//https://dev.to/codypotter/layered-architectures-in-go-3cg8
package main

import (
	"fmt"

	"github.com/juanheza/facturador/httplayer"
)

func main() {
	fmt.Println("Initializing")
	httplayer.New()
	fmt.Println("Server Started")
}

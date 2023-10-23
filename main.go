package main

import (
	"fmt"
	"helloworld/github.com/frsarmientop/godesde0/variables"
)

func main() {
	estado, texto := variables.ConviertoaTexto(1500)
	fmt.Println(estado)
	fmt.Println(texto)
}

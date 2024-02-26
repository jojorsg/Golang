package main

import (
	"introducao-testes/enderecos"
	"fmt"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Rodovia dos Paulista")
	fmt.Println(tipoEndereco)
}
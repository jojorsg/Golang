package main

import (
	"fmt"
	"modulo/auxiliar"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("josueroberto099@gmail.com")
	fmt.Println(erro)

	erro2 := checkmail.ValidateFormat("741")
	fmt.Println(erro2)
}

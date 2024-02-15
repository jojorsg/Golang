package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero uint32
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Josué"
	u.idade = 24
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos Lobos", 1307}

	usuario2 := usuario {"Rodrigo", 24, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Rebeca"}
	fmt.Println(usuario3)

	usuario4 := usuario{idade: 22}
	fmt.Println(usuario4)
}
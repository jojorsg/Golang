package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	idade uint8
	altura float32
}

type estudante struct {
	pessoa
	curso string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa {"João", "Pedro", 20, 1.78}
	fmt.Println(p1)

	e1 := estudante {p1, "Engenharia", "UFC"}
	fmt.Println(e1)
}
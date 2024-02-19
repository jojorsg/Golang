package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}
func main() {
	usuario1 := usuario {"Usuário 1", 20}
	usuario1.salvar()

	usuario2 := usuario {"Jojo", 29}
	usuario2.salvar()
	maiorIdade := usuario2.maiorDeIdade()
	fmt.Println(maiorIdade)
}

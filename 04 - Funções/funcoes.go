package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func ()  {
		fmt.Println("Função F")
	}

	f()

	var f1 = func (txt string)  {
		fmt.Println(txt)
	}

	f1("Texto de f1")

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	// COM FUNÇÕES COM MAIS DE UM RETORNO, PODEMOS SELECIONAR APENAS UM RETORNO UTILIZANDO O _
	resultadoSoma1, _ := calculosMatematicos(5, 7)
	fmt.Println(resultadoSoma1)

	// A ORDEM DOS RETORNOS NAS FUNÇÕES SÃO IMPORTANTES E DEVEM SER RESPEITADAS
	_ , resultadoSubtracao1 := calculosMatematicos(26, 13)
	fmt.Println(resultadoSubtracao1)
}
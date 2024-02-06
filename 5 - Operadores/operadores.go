package main

import "fmt"

func main()  {
	// ARITMÉTICOS
	soma := 1 + 2
	subtracao := 1 - 2
	multiplicacao := 10 * 5
	divisao := 10 / 4
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, multiplicacao, divisao, restoDaDivisao)

	//ATRIBUIÇÃO
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	//OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	fmt.Println("*-----------------------*")

	//OPERADORES LÓGICOS
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	fmt.Println("*-----------------------*")

	//OPERADORES UNÁRIOS
	numero := 10
	numero++
	numero += 15
	numero--
	numero -= 10
	fmt.Println(numero)

	numero *= 3
	numero /= 10
	numero %= 3
	fmt.Println(numero)
}
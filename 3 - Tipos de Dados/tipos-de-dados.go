package main

import "fmt"

func main() {
	
	// NÚMEROS INTEIROS

	numero := 1000000000000000000
	fmt.Println(numero)

	var numero1 int64 = -1000000000000000000
	fmt.Println(numero1)

	var numero2 uint32 = 1000000000
	fmt.Println(numero2)

	// alias = apelido
	// INT32 = RUNE
	var numero3 rune = 1234567890
	fmt.Println(numero3)

	// BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)


	// NÚMEROS REAIS

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123456789.00000
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)


	// STRINGS

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto02"
	fmt.Println(str2)

	char := 'J'
	fmt.Println(char)
}
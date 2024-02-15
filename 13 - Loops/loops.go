package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i<10 {
		i++
		fmt.Println(i)
		time.Sleep(time.Second)
	}

	for j:= 10; j>=1; j-- {
		fmt.Println(j)
		time.Sleep(time.Second)
	}

	nomes := [3]string {"Jojo", "Rodrigo", "Milena"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, letra)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string {
		"nome": "Jojo",
		"sobrenome": "Alves",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}
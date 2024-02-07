package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string {
		"nome": "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario)

	usuario2 := map[string]map[string]string {
		"nome": {
			"primeiro": "João",
			"último": "Soares",
		},
		"curso": {
			"nome": "Engenharia",
			"campus": "Campus 1",
		},
	}

	fmt.Println(usuario2)

	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["sexo"] = map[string]string {
		"sexo": "masculino",
	}

	fmt.Println(usuario2)
}
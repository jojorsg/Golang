package main

import "fmt"

func main() {
	fmt.Println("Arrays e Slices")

	var array1[5] string
	array1[2] = "Posição 3"
	fmt.Println(array1)

	array2 := [5]string {"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	array3 := [...]int {1, 2, 3, 4, 5, 265 }
	fmt.Println(array3)

	slice := []int{10,11,12,13,14,15,16,18,18,19,25,456,996,6699}
	fmt.Println(slice)

	slice = append(slice, 18, 19)
	fmt.Println(slice)

	slice2 := array2[2:4]
	fmt.Println(slice2)

	array2[3] = "Alteração"
	fmt.Println(slice2)

	// ARRAYS INTERNOS
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacidade

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacidade

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4)) // length
	fmt.Println(cap(slice4)) // capacidade
}
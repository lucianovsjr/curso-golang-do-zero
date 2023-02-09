package main

import (
	"fmt"
	"reflect"
)

/*
	Arrays e slices são estruturas parecidas.

	Array: Lista de valores de apenas um tipo e tamanho fixo.
	Slice: Estrutura baseada no Array com mais flexibilidade. Não possui tamanho fixo.

	Slice é mais usado do que o Array.

	Slice é uma fatia de um array. Aponta para um array.
*/

func main() {
	var array1 [5]string // valor zero nas posições
	array1[0] = "Position 1"
	fmt.Println(array1)

	array2 := [5]string{"Position 1", "Position 2", "Position 3", "Position 4", "Position 5"}
	array3 := [5]string{}
	array4 := [5]string{"Position 1"}
	fmt.Println(array2)
	fmt.Println(array3)
	fmt.Println(array4)

	// ... fixa o tamanho do array de acordo com os valores definidos
	array5 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array5)

	slice1 := []int{}
	slice2 := []int{10, 11}
	fmt.Println(slice1, slice2)

	// Array e Slice são tipos diferentes
	fmt.Println(reflect.TypeOf(array5))
	fmt.Println(reflect.TypeOf(slice2))

	slice2 = append(slice2, 12) // Retorna um slice novo
	fmt.Println(slice2)

	// criar um slice apartir de um array. Aponta para aquela fatia do array.
	slice3 := array2[1:3] // param1 inclusivo e param2 exclusivo. >= e <
	fmt.Println(slice3)

	array2[1] = "Changed position"
	fmt.Println(array2)
	fmt.Println(slice3)
}

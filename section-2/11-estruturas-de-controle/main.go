package main

import "fmt"

func main() {
	num1 := 10

	// No Go não é usado os parênteses na condição.
	// A própria extensão do Go remove os parênteses.
	// Parênteses pode ser usado em operações que possuem uma precedência especifica.
	// Só existe uma forma de estruturar um if.
	if num1 > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor que 15")
	}

	// if init, quando inicializamos uma variável na condição do if
	// escopo da variável é somente dentro do if
	if otherNumber := num1; otherNumber > 10 {
		fmt.Println("Number is greater than ten", otherNumber)
	} else if otherNumber2 := otherNumber; otherNumber2 > 0 {
		fmt.Println("Number is greate than zero", otherNumber, otherNumber2)
	} else {
		fmt.Println("Number is zero", otherNumber, otherNumber2)
	}
}

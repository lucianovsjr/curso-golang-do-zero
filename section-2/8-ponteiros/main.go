package main

import "fmt"

/*
	Ponteiro é uma referência de memória.
	Bastante utilizado pelo Go.
*/

func main() {
	// Atribuindo uma variável a outra:
	// Na forma padrão é feito uma copia do valor
	// No ponteiro é copiado a referência de memória
	var num1 int = 10
	var num2 int = num1

	fmt.Println(num1, num2)

	num1++
	fmt.Println(num1, num2)

	// Ponteiro
	var num3 int = 100
	var pointer1 *int = &num3

	// Valor e o endereço de pointer1 e num3 são iguais
	fmt.Println(num3, *pointer1) // desreferenciação: * pega o valor salvo em um endereço
	fmt.Println(&num3, pointer1) // & pega o endereço de uma variável

	num3 = 150
	fmt.Println(num3, *pointer1)
	fmt.Println(&num3, pointer1)
}

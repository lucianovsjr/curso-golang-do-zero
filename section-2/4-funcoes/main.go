package main

import "fmt"

/*
	Função é um tipo.

	É possível passar uma função como parâmetro para uma função.
	É possível retornar uma função de uma função.
*/

func main() {
	sum := somar(10, 20)
	fmt.Println(sum)

	var f = func(text string) bool {
		fmt.Println("Function F:", text)
		return true
	}
	var ret = f("Texto")
	fmt.Println("Return F:", ret)

	// Go te obriga utilizar todos os retornos e não é possível declarar uma var e não utilizar
	// Para omitir retornos, use "_"
	calcResult, _, _ := mathCalc(1, 1, "abc")
	fmt.Println(calcResult)
}

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

/*
É possível omitir o tipo de um parametro para pegar o tipo do proximo parametro.
É possível ter vários retornos na função.
*/
func mathCalc(n1, n2 int8, t string) (int8, int8, string) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub, t
}

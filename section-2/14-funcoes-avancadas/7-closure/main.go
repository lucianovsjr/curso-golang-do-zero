package main

import "fmt"

/*
	closure:
	Funções que referenciam variáveis que estão fora do seu corpo.
*/

func closure() func() {
	text := "Inside of the clouser function"

	func1 := func() {
		// A referência de text é mantida
		fmt.Println(text)
	}

	return func1
}

func main() {
	text := "Inside of the main function"
	fmt.Println(text)

	newFunc := closure()
	newFunc()
}

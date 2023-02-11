package main

import "fmt"

/*
	Função anônima: Basicamente uma função sem nome.
*/

func main() {
	// No exemplo abaixo, assim que o Go termina de declar a função anônima, ele a executa.
	func(text string) {
		fmt.Println(text)
	}("Hello")

	// any: interface vazia, utiliza como parâmetro variático na função Sprintf
	text := func(name string, value any) string {
		// Sprintf concatena string
		return fmt.Sprintf("Hello %s you is %d", name, value)
	}("Jr", 10)
	fmt.Println(text)
}

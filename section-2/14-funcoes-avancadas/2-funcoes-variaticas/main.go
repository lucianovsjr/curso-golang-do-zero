package main

import "fmt"

/*
	Função variática: Função que pode receber N parametros

	Uso dos "..." identifica os N parametros, a variável recebe um slice.

	Se nenhum parâmetro for passado, a variável recebe um slice vazio.
*/

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

// A função pode ter apenas um parâmetro variático e ele deve ser o último.
func writter(text string, numbers ...int) {
	for _, number := range numbers {
		fmt.Println(text, number)
	}
}

func main() {
	sumTotal := sum(1, 2, 3)
	fmt.Println(sumTotal)

	writter("Hello", 2, 4, 6, 8)
}

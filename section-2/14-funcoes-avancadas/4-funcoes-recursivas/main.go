package main

import "fmt"

/*
	Função Recursiva:
	Depende de uma execução dela própria para terminar a execução.
	Precisa especificar o momento que irá parar de chamar a si mesmo,
	para não acontecer estouro de pila (stackoverflow).

	Não é recomendado em casos que irá acontecer muitas execuções recursivas.
*/

// Sequência de fibonacci: O próximo número é a soma dos 2 números anteriores
// 1 1 2 3 5...
func fibonacci(pos uint) uint {
	// condição de parada
	if pos <= 1 {
		return pos
	}

	return fibonacci(pos-2) + fibonacci(pos-1)
}

func main() {
	pos := uint(8)
	fmt.Println(fibonacci(pos))

	for i := uint(1); i <= pos; i++ {
		fmt.Println(fibonacci(i))
	}

}

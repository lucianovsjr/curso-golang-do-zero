package main

import "fmt"

/*
	Função init:
	É a função que irá executar antes da função main.
	Podemos ter uma por arquivo.
	Se tivermos um pacote com 10 arquivos, cada um, pode ter a função init.
	Pode ser utilizada para setup, inicialização de variáveis...
*/

var n int

func init() {
	fmt.Println("Init", n)
	n = 100
}

func main() {
	fmt.Println("Main", n)
}

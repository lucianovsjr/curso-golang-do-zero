package main

import "fmt"

// Retorno nomeado
// Irá retornar sum e minus
// return já sabe quais variáveis precisa retornar
func calc(num1, num2 int) (sum, minus int) {
	sum = num1 + num2
	minus = num1 - num2
	return
}

// retorno nomeado não é muito utilizado mas deixa o código mais legível
func main() {
	sum, minus := calc(1, 1)
	fmt.Println("Calc", sum, minus)
}

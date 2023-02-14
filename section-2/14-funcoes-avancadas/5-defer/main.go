package main

import "fmt"

/*
	defer, significa adiar.
	Adia a execução de algo até o último momento possível.

	Quanto utilizado com uma func que possui return, irá executar o defer
	antes de qualquer return.

	Exemplo: Fechar conexões com bd em caso de sucesso e falha.
*/

func func1() {
	fmt.Println("func1")
}

func func2() {
	fmt.Println("func2")
}

func studentIsApproved(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado será retornado")     // last
	defer fmt.Println("Média calculada 2. Resultado será retornado 2") // second
	fmt.Println("studentIsApproved")

	avg := (n1 + n2) / 2
	defer fmt.Println("Média calculada 3. Resultado será retornado 3") // first

	if avg >= 6 {
		return true
	}

	return false
}

func main() {
	defer func1()
	func2()

	fmt.Println(studentIsApproved(6, 8))
}

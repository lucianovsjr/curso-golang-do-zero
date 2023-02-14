package main

import "fmt"

/*
	panic:
	Interrompe o fluxo normal do programa.
	Chama todas os defer.
	Nada depois será executado.
	Não é a melhor maneira de tratar errors.
	Não é um erro.
	Como se quebrasse o programa.

	recover:
	Recupera o panic gerado no programa com o intuito de tratar o panic.
*/

func recoverExec() {
	// Se não estiver entrando em panic, err será nil
	if err := recover(); err != nil {
		fmt.Println("Tentativa de recurperar", err)
	}
}

func studentIsApproved(n1, n2 float64) bool {
	defer recoverExec()
	avg := (n1 + n2) / 2

	if avg > 6 {
		return true
	} else if avg < 6 {
		return false
	}

	panic("Average is exactly 6!")
}

func main() {
	fmt.Println(studentIsApproved(6, 6))
}

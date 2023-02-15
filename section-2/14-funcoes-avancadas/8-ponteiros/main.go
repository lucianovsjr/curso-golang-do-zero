package main

import "fmt"

func signalReverse(num int) int {
	return num * -1
}

// num é um ponteiro do tipo int
func signalReverseWithPointer(num *int) {
	fmt.Println(num)
	// "*" desreferenciação
	*num = *num * -1
}

func main() {
	num1 := 20
	reversedNum := signalReverse(num1) // Passado o valor
	fmt.Println(reversedNum)
	fmt.Println(num1)

	num2 := 40
	fmt.Println(num2)
	signalReverseWithPointer(&num2) // "&" pega o endereço de memória para passar
	fmt.Println(num2)
}

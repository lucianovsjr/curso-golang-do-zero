package main

import "fmt"

func main() {
	// Aritméticos
	sum := 1 + 2
	minus := 1 - 2
	divide := 1 / 2
	multiply := 1 * 2
	modulo := 1 % 2

	fmt.Println(sum, minus, divide, multiply, modulo)

	// int16, int32... são tipos diferentes. Não é possível somar...
	var num1 int16 = 10
	// var num2 int32 = 25
	var num2 int16 = 25
	sum2 := num1 + num2
	fmt.Println(sum2)

	// Atribuição
	var var1 string = "String"
	var2 := "String"
	fmt.Println(var1, var2)

	// Relacionais
	fmt.Println("1 > 2", 1 > 2)
	fmt.Println("1 < 2", 1 < 2)
	fmt.Println("1 <= 2", 1 <= 2)
	fmt.Println("1 >= 2", 1 >= 2)
	fmt.Println("1 == 2", 1 == 2)
	fmt.Println("1 != 2", 1 != 2)

	// Lógicos
	// Não existe o xor (ou exclusivo)
	fmt.Println("true && false", true && false)
	fmt.Println("true || false", true || false)
	fmt.Println("!true", !true)

	// Unários
	num6 := 10
	num6++ // Só fuciona o pos-fxiado, pré-fixado não existe
	num6 += 15
	fmt.Println(num6)

	num6-- // Só fuciona o pos-fxiado, pré-fixado não existe
	num6 -= 15
	fmt.Println(num6)

	num6 *= 2
	fmt.Println(num6)

	num6 /= 2
	fmt.Println(num6)

	num6 %= 3
	fmt.Println(num6)

	// Ternário não existe
	// Premissa do GO é existir apenas uma forma de fazer as coisas
	var texto string
	if num6 > 0 {
		texto = "Maior que 0"
	} else {
		texto = "Não é maior que 0"
	}
	fmt.Println(texto)
}

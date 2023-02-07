package main

import (
	"errors"
	"fmt"
)

// O compilador coloca o ponto e vírgula mas ele não é utilizado no desenvolvimento do código.
// se colocar o ponto e vírgula a extensão irá remover.
// Funções também são tipos

func main() {
	int_type()
	float_type()
	string_type()
	bool_type()
	error_type()
	zero_value()
}

/*
	Tipos:
		int8 	-128 to 127
		int16   -32768 to 32767
		int32   -2147483648 to 2147483647
		int64   -9223372036854775808 to 9223372036854775807
		int     Depends on platform: 32 bits in 32 bit systems and 64 bit in 64 bit systems
*/

func int_type() {
	fmt.Println("--------int_type--------")
	var num int16 = 100
	fmt.Println(num)

	// int utilizar o bit da arquitetura do seu computador
	var num2 int = 1000
	fmt.Println(num2)

	// Infere int
	num3 := 1000
	fmt.Println(num3)

	// unsygned int, número sem sinal
	// unit8, unit16, unit32, unit64
	var num4 uint = 32
	fmt.Println(num4)

	// alias, exatamente igual ao int32
	// Utilizado quando está trabalhando com número que representam caracteres,
	// principalmente da tabela asc
	var num5 rune = 123456
	fmt.Println(num5)

	// alias, exatamete igual ao uint8
	var num6 byte = 1
	fmt.Println(num6)
}

// tipos: float32, float 64
func float_type() {
	fmt.Println("--------float_type--------")
	var num float32 = 100.5
	fmt.Println(num)

	var num2 float64 = 10011111.55
	fmt.Println(num2)

	// Infere float de acordo com a arquitetura do seu computador
	num3 := 12345.67
	fmt.Println(num3)
}

func string_type() {
	fmt.Println("--------string_type--------")
	var str string = "AAABBBCCC"
	fmt.Println(str)

	str2 := "DDDEEEFFF"
	fmt.Println(str2)

	// No Go não existe o tipo char
	// Algo similar
	// Com aspas simples armazena o número do caracter da tabela asc
	char := 'B'
	fmt.Println(char)
}

func bool_type() {
	fmt.Println("--------bool_type--------")
	var bool1 bool = true
	fmt.Println(bool1)
}

func error_type() {
	fmt.Println("--------error_type--------")
	var erro error = errors.New("Internal error")
	fmt.Println(erro)
}

// Valor zero: Valor atribuido a uma variável quando ela não recebe um valor
func zero_value() {
	fmt.Println("--------zero_value--------")
	// string vazia, ""
	var texto string
	fmt.Println(texto)

	// Valor zero, 0
	var num int
	fmt.Println(num)

	// Valor zero, 0
	var num2 float64
	fmt.Println(num2)

	// false
	var boolean bool
	fmt.Println(boolean)

	// nil
	var erro error
	fmt.Println(erro)
}

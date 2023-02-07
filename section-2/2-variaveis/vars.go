package main

import "fmt"

/*
2 Formas de declarar variáveis:
	1. Deixando o valor implícito
	2. Deixando o valor explícito

Go não deixa criar uma variável e não usar, o compilador gera um error.
	ex: "var2 declared and not used"
O mesmo acontece para funções, imports...

No Go a maioria das coisas tem apenas um jeito de ser feito mas variavel não é uma delas.
*/

func main() {
	// explícito
	var var1 string = "var 1"
	fmt.Println(var1)

	// implícito (inferência de tipo)
	var2 := "var 2"
	fmt.Println(var2)

	// declarar mais de uma variável
	var (
		var3 string = "var 3"
		var4 string
	)
	var4 = "var 4"
	fmt.Println(var3, var4)

	var5, var6 := "var 5", "var 6"
	fmt.Println(var5, var6)

	const const1 string = "const 1"
	fmt.Println(const1)

	const (
		const2 string = "const 2"
		const3 string = "const 3"
	)
	fmt.Println(const2, const3)

	var5, var6 = var6, var5
	fmt.Println(var5, var6)
}

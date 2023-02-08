package main

import "fmt"

/*
	Structs é o mais próximo que o GO tem das classes.
	No GO não temos classes.

	Struct é uma coleção de campos.

	Valor zero de uma struct é a coleção de campos com o valor zero.
*/

type user struct {
	name    string
	age     uint8
	address address
}

type address struct {
	street string
	number uint8
}

func main() {
	var user1 user
	fmt.Println(user1)
	user1.name = "Jr"
	user1.age = 18
	fmt.Println(user1)

	// Inferencia
	address1 := address{"Rua 1", 22}
	user2 := user{"Lj", 25, address1}
	fmt.Println(user2)

	// Não preenchendo todos os campos
	// user3 := user{"Jl"} => Error
	user3 := user{name: "Jl"}
	fmt.Println(user3)

	user4 := user{age: 1}
	fmt.Println(user4)
}

package main

import "fmt"

/*
	No Go não existe herança.
	Temos algo que podemos assimilar com a herança, mas não é.

	Podemos criar as structs e não usa-lás.
*/

type person struct {
	name     string
	age      uint8
	height   uint8
	lastName string
}

// student irá ter os campos de person
// não é passado o tipo
// os campos de person podem ser acessados como se fossem de student
type student struct {
	person
	course  string
	college string
}

func main() {
	person1 := person{"João", 21, 178, "Silva"}
	fmt.Println(person1)

	student1 := student{person1, "Engenharia", "Estacio"}
	fmt.Println(student1)

	// Não precisamos acessar person para pegar um campo que pertence a essa struct
	fmt.Println(student1.person.name)
	fmt.Println(student1.name)
}

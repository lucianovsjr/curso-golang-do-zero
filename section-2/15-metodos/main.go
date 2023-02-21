package main

import "fmt"

/*
	Método:
	É uma função que está associada a algo, estruct ou interface por exemplo.
*/

type user struct {
	name string
	age  uint8
}

// funcao save associada a struct user
// Não pode estar associada a mais de uma estrutura
func (u user) save() {
	fmt.Printf("Salvando os dados do usuário \"%s\" no banco de dados\n", u.name)
}

func (u user) ofLegalAge() bool {
	return u.age >= 18
}

// Alterando um dado do struct
// ponteiro de user, é recebio uma referência da struct e não uma copia.
func (u *user) doBirth() {
	u.age++
}

func main() {
	user1 := user{"User 1", 20}
	fmt.Println(user1)
	user1.save()

	user2 := user{"Jhon", 22}
	fmt.Println(user2)
	user2.save()

	ofLegalAge2 := user2.ofLegalAge()
	fmt.Println(ofLegalAge2)

	user2.doBirth()
	fmt.Println(user2.name, user2.age)
}

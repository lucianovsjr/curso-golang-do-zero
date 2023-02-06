package main

import (
	"fmt"
	"pacotes/auxi"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Main do módulo pacotes")
	auxi.Writer()

	email := "abc#hotmail.com"
	error := checkmail.ValidateFormat(email)
	fmt.Println(email, error)

	email2 := "abc@hotmail.com"
	error2 := checkmail.ValidateFormat(email2)
	fmt.Println(email2, error2)
}

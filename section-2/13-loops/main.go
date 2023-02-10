package main

import (
	"fmt"
	"time"
)

/*
 Go só tem o For como estrutura de repetição
*/

func main() {
	// Parecido com o while
	i := 0
	for i < 3 {
		fmt.Println("i", i)
		time.Sleep(time.Second)
		i++
	}

	// for padrão
	for j := 0; j < 4; j++ {
		fmt.Println("j", j)
		time.Sleep(time.Second / 2)
	}

	for j := 0; j < 10; j += 2 {
		fmt.Println("j", j)
		time.Sleep(time.Second / 2)
	}

	// range, itera array, slice, string e map
	names := [3]string{"João", "Joaquin", "Jão"}
	for index, name := range names {
		fmt.Println(index, name)
	}

	// omitir o indice
	for _, name := range names {
		fmt.Println(name)
	}

	// Traz os números da tabela asc para cada letra
	for index, letter := range "STRING" {
		fmt.Println(index, letter, string(letter))
	}

	users := map[string]string{
		"name": "Jr",
		"age":  "18",
	}
	for key, value := range users {
		fmt.Println(key, value)
	}

	// range para struct não funciona

	// loop infinto
	// Parecido com o for true {} de outras linguagens
	for {
		fmt.Println("Infinite loop")
		break // Sair do loop
	}

	// continue para pular para a proxima iteração
}

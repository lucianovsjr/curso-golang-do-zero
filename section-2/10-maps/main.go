package main

import "fmt"

/*
	Map, estrutura de chave e valor.
	A chave e o valor tem sempre um tipo especifico.
*/

func main() {
	// map["tipo da key"]"tipo do value"
	user := map[string]string{
		"name":     "Jr",
		"lastName": "Silva",
	}
	fmt.Println(user)
	fmt.Println(user["name"], user["lastName"]) // Para acessa não pode usar a notação ponto

	item := map[int8]string{
		1: "Item 1",
		2: "Item 2",
	}
	fmt.Println(item)
	fmt.Println(item[1], item[2])

	// Map aninhado
	// Muitos maps aninhados fica dificil de entender
	user2 := map[string]map[string]string{
		"name": {
			"first": "Jr",
			"last":  "Silva",
		},
		"name2": {
			"first": "Jr",
			"last":  "Silva",
		},
	}
	fmt.Println(user2)
	fmt.Println(user2["name"]["first"], user2["name"]["last"])

	delete(user2, "name2")
	fmt.Println(user2)

	// Add key
	user2["name3"] = map[string]string{
		"full": "Jr Silva",
	}
	fmt.Println(user2)
}

package main

import "fmt"

/*
	package main:
		Dizer para o go que esse arquivo está dentro do pacote main.
		Obrigatóriamente para ser executavel.
	package:
		Go organizar os arquivos em pacotes.
		Grupo de arquivos que estão no mesmo diretorio, são compilados juntos.
		Existe um compartilhamente de funções, variáveis... que estão no mesmo pacote.
	import:
		Referência dos pacotes que serão utilizados no arquivo.
	func main:
		entry point.

	Para executar esse arquivo: go run first.go
*/

func main() {
	fmt.Println("Hello world!")
}

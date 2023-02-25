package main

import (
	"fmt"
	"math"
)

/*
	Interfaces:
	Muito utilizada quando é necessário ter uma certa flexibilidade com os tipos.
*/

// Possui apenas assinaturas de métodos
type shape interface {
	area() float64
}

type retangulo struct {
	height float64
	width  float64
}

func (r retangulo) area() float64 {
	return r.height * r.width
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

// o parametro passado para a função deve satisfazer a assinatura da interface shape.
// Verifica se o parametro passado possui os métodos da interface.
// Implementação da interface é implicita
func areaWritter(s shape) {
	fmt.Printf("A área da forma é %0.2f\n", s.area())
}

func main() {
	r := retangulo{10, 15}
	areaWritter(r)

	c := circulo{10}
	areaWritter(c)
}

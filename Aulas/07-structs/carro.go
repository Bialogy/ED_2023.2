package main

import "fmt"

type Carro struct {
	Modelo     string
	Cor        string
	Velocidade int
	Ligado     bool
}

func (c *Carro) Ligar() {
	c.Ligado = !c.Ligado
}

func (c *Carro) Acelerar(valor int) {
	c.Velocidade += valor
}

func main() {
	carro = Carro{
		Modelo:     "gol",
		Cor:        "preto",
		Velocidade: 0,
		Ligado:     false,
	}

	fmt.Println(carro)
	carro.Ligar()
	fmt.Println(carro)
	carro.Acelerar(10)
	fmt.Println(carro)
}

package cli

import (
	"fmt"
	p "mcronalds/produtos"
)

func Cli() {
	var prod1, prod2 p.Produto
	prod1 = p.Produto{ Nome: "produto 1" }
	prod2 = p.Produto{ Nome: "produto 2" }

	prod1.DefinirId()
	prod2.DefinirId()

	fmt.Println(p.TotalProdutos)
	fmt.Println(prod1)
	fmt.Println(prod2)
}
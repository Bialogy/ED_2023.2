package main

import (
	"fmt"
	ut "projeto/utils" // import utils as ut
)

func main() {
	fmt.Println("Olá, mundo!")

	fmt.Println(ut.Somar(6.4, 8.2))
	fmt.Println(ut.Multiplicar(6.4, 8.2))
}

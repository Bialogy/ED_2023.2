package main

import "fmt"

func main() {
	var x, y, z = 4, 10, -2

	// aritméticos
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(x % y)

	// atribuição
	z += 2 // z = z + 2
	z -= x
	z *= 3
	z /= 2
	z %= 3

	// comparação
	fmt.Println(x == y)
	fmt.Println(x != y)
	fmt.Println(x < y)
	fmt.Println(x <= y)
	fmt.Println(x > y)
	fmt.Println(x >= y)

	// lógicos
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	// não vamos ver agora
	// bitwise
	// memória e canal
}
package main

import "fmt"

const M = 3

func main() {
	var fila [M]int
	in, fim = -1, -1

	insereFila(&fila, &in, &fim, 4)
	fmt.Println(fila, in, fim)
	insereFila(&fila, &in, &fim, 5)
	fmt.Println(fila, in, fim)
	removeFila(&fila, &in, &fim)
	fmt.Println(fila, in, fim)
	insereFila(&fila, &in, &fim, 2)
	fmt.Println(fila, in, fim)
	insereFila(&fila, &in, &fim, 1)
	fmt.Println(fila, in, fim)
	insereFila(&fila, &in, &fim, 6)
	fmt.Println(fila, in, fim)
	removeFila(&fila, &in, &fim)
	fmt.Println(fila, in, fim)
	removeFila(&fila, &in, &fim)
	fmt.Println(fila, in, fim)
	insereFila(&fila, &in, &fim, 6)
	fmt.Println(fila, in, fim)
	removeFila(&fila, &in, &fim)
	fmt.Println(fila, in, fim)
	removeFila(&fila, &in, &fim)
	fmt.Println(fila, in, fim)
	removeFila(&fila, &in, &fim) // Underflow
	fmt.Println(fila, in, fim)

	// var pilha [M]int
	// topo = 0

	// inserePilha(&pilha, &topo, 4)
	// fmt.Println(pilha)
	// inserePilha(&pilha, &topo, 5)
	// fmt.Println(pilha)
	// inserePilha(&pilha, &topo, 2)
	// fmt.Println(pilha)
	// inserePilha(&pilha, &topo, 1) // Overflow
	// fmt.Println(pilha)

	// fmt.Println(removePilha(&pilha, &topo))
	// fmt.Println(pilha)
	// fmt.Println(removePilha(&pilha, &topo))
	// fmt.Println(pilha)
	// fmt.Println(removePilha(&pilha, &topo))
	// fmt.Println(pilha)
	// fmt.Println(removePilha(&pilha, &topo)) // Underflow, -1
	// fmt.Println(pilha)
}

func removeFila(f *[M]int, in *int, fim *int) {
	if *in == *fim && *fim == -1 {
		fmt.Println("Underflow")
		return
	}

	f[*in] = 0
	if *in == *fim {
		*in, *fim = -1, -1
	} else {
		*in = (*in + 1) % M
	}
}

func insereFila(f *[M]int, in *int, fim *int, valor int) {
	if (*fim+1)%M == *in {
		fmt.Println("Overflow")
		return
	}

	if *in == *fim && *fim == -1 {
		*in, *fim = 0, 0
	} else {
		*fim = (*fim + 1) % M
	}

	f[*fim] = valor
}

func inserePilha(p *[M]int, topo *int, valor int) {
	if *topo == M {
		fmt.Println("Overflow")
		return
	}

	p[*topo] = valor
	*topo++
}

func removePilha(p *[M]int, topo *int) int {
	if *topo == 0 {
		fmt.Println("Underflow")
		return -1
	}

	*topo--
	valorRemovido = p[*topo]
	p[*topo] = 0
	return valorRemovido
}

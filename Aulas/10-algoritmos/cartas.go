package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	n1, n2 = 0, 0      // eles poderiam ser int8
	classificacao = "" // poderia ser do tipo rune

	// Lê uma linha inserida pelo usuário
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()

	linha = scanner.Text()

	// Divide a linha em elementos, utilizando espaço como separador
	// numeros é um slice de strings
	numeros = strings.Split(linha, " ")

	for _, strNumero = range numeros {
		// O segundo retorno é uma mensagem de erro, caso dê algum problema
		// Estamos assumindo que não vai dar problema
		n2, _ = strconv.Atoi(strNumero)

		if n1 == 0 {
			n1 = n2
			continue
		}

		if n2 > n1 {
			if classificacao == "" {
				classificacao = "C"
			} else if classificacao == "D" {
				classificacao = "N"
				break
			}
		}

		if n1 > n2 {
			if classificacao == "" {
				classificacao = "D"
			} else if classificacao == "C" {
				classificacao = "N"
				break
			}
		}

		n1 = n2
	}

	fmt.Println(classificacao)
}

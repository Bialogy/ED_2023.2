package main

import "fmt"

func main() {
	var expressoes []string = make([]string, 0)

	for {
		var expressao string
		fmt.Scanln(&expressao)

		if expressao == "" {
			break
		}

		expressoes = append(expressoes, expressao)
	}

	for _, expressao = range expressoes {
		fmt.Println(encontreOTelefone(expressao))
	}
}

func encontreOTelefone(expr string) string {
	telefone = ""

	for _, carac = range expr {
		switch carac {
		case '1', '0', '-':
			telefone += string(carac)
		case 'A', 'B', 'C':
			telefone += "2"
		case 'D', 'E', 'F':
			telefone += "3"
		case 'G', 'H', 'I':
			telefone += "4"
		case 'J', 'K', 'L':
			telefone += "5"
		case 'M', 'N', 'O':
			telefone += "6"
		case 'P', 'Q', 'R', 'S':
			telefone += "7"
		case 'T', 'U', 'V':
			telefone += "8"
		case 'W', 'X', 'Y', 'Z':
			telefone += "9"
		}
	}

	return telefone
}

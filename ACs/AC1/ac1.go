func main() {
	var primo int
	fmt.Println("Escreva um número para checar se é primo: ")
	fmt.Scanln(&primo)

	var fibo int
	fmt.Println("Escreva um número para achar seu termo na sequência de Fibonacci: ")
	fmt.Scanln(&fibo)

	var semana int
	fmt.Println ("Escreva um número para verificar seu dia da semana: ")
	fmt.Scanln(&semana)

	var bissexto int
	fmt.Println("Escreva um ano para verificar se é bissexto: ")
	fmt.Scanln(&bissexto)
}

func e_primo(primo) {

	if primo <= 1 {
		fmt.Println("Não é primo.")
	} else for i := 2; i < primo; i++ {
				if primo % i {
					fmt.Println(i)
				}
		}

}

func fibo(fibo) {

	if fibo <= 0 {
		return 0
	} else if fibo == 1 {
		return 1
	} else {
		return fibo(n -1) + fibo(n-2)
	}

}

func dia_semana(semana) {

	if semana == 1 {
		return "1-Domingo"
	} else if semana == 2 {
		return "2-Segunda"
	} else if semana == 3 {
		return "3-Terca"
	} else if semana == 4 {
		return "4-Quarta"
	} else if semana == 5 {
		return "5-Quinta"
	} else if semana == 6 {
		return "6-Sexta"
	} else if semana == 7 {
		return "7-Sabado"
	} else {
		return "dia da semana inválido."
	}

}

func e_bissexto(bissexto) {

	if (bissexto % 4 == 0 && e_bissexto % 100 != 0) || (bissexto % 400 == 0) {
		return "é bissexto."
	} else {
		return "não é bissexto."
	}
}
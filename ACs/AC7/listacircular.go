package main

import "fmt"

type No struct {
	valor int
	prox  *No
}

type Lista struct {
	ini *No
	fim *No
	len int
}

func (l *Lista) insere(valor int) {
	novoNo := &No{valor: valor}

	novoNo.prox = l.ini
	l.ini = novoNo
	l.fim.prox = novoNo
	
}

func (l *Lista) remove() {
	l.ini = l.ini.prox
	l.fim.prox = l.ini
}

func (l *Lista) exibe() {
	noAtual := l.ini
	
	for i := 0; i < l.len; i++ {
		fmt.Println(noAtual.valor)
		noAtual = noAtual.prox
	}
}

func main() {
	l := *&Lista{}

	l.insere(25)
	l.insere(20)
	l.insere(15)
	l.insere(10)
	l.insere(5)

	fmt.Println("Lista:")
	l.exibe()
	
	l.remove()

	fmt.Println("Lista sem o primeiro nó:")
	l.exibe()

}


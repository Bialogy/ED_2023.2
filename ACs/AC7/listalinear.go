package main

import "fmt"

type No struct {
	valor int
	prox *No
	ant *No
}

type Lista struct {
	ini *No
	fim *No
}

func (l *Lista) busca(chave int) *No {
	
	noAtual := l.ini
	for noAtual != nil {
		if noAtual.valor == chave {
			return noAtual
		}
		noAtual = noAtual.prox
	}

	return nil

}

func (l *Lista) insere(posNo *No, valor int){

	novoNo := &No{valor: valor}

	if posNo == nil {
		return
	}

	novoNo.prox = posNo.prox
	novoNo.ant = posNo

	if posNo.prox != nil {
		posNo.prox.ant = novoNo
	} else {
		l.fim = novoNo
	}

	posNo.prox = novoNo

}

func (l *Lista) remove(posNo *No) {

	if posNo == nil {
		return
	}

	if posNo == l.ini{
		l.ini = posNo.prox
	}

	if posNo == l.fim {
		l.fim = posNo.ant
	}

	if posNo.ant != nil {
		posNo.ant.prox = posNo.prox
	}

	if posNo.prox != nil {
		posNo.prox.ant = posNo.ant
	}

}

func (l *Lista) exibe() {
	noAtual := l.ini

	for noAtual != nil {
		fmt.Println(noAtual.valor)
		noAtual = noAtual.prox
	}
}

func main() {

	l := &Lista{}

	l.insere(nil, 5)
	l.insere(nil, 10)
	l.insere(nil, 15)
	l.insere(nil, 20)
	l.insere(nil, 25)

	fmt.Println("Lista: ")
	l.exibe()


}

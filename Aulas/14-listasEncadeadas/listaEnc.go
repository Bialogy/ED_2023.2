package main

import "fmt"

type No struct {
	chave int
	prox  *No
}

type Lista struct {
	cab *No
}

// exibição da lista
func (l *Lista) exibe() {
	no = l.cab

	for no != nil {
		fmt.Println(no)
		no = no.prox
	}
}

// busca simples
func (l *Lista) buscaSimples(chave int) *No {
	no = l.cab

	for no != nil {
		if no.chave == chave {
			return no
		}
		no = no.prox
	}

	return nil
}

// inserção simples, ao final da lista
func (l *Lista) insere(ch int) {
	novoNo = &No{chave: ch}

	no = l.cab
	if no == nil {
		l.cab = novoNo
	} else {
		for no.prox != nil {
			no = no.prox
		}
		no.prox = novoNo
	}
}

// considerando uma lista ordenada:
// busca de uma lista ordenada
func (l *Lista) busca(ch int) (*No, *No) {
	var ant *No
	no = l.cab

	if no == nil {
		return ant, no
	}
	for no.chave < ch {
		ant = no
		no = no.prox
		if no == nil {
			return ant, no
		}
	}
	if no.chave == ch {
		return ant, no
	}

	return ant, nil
}

// inserção de uma lista ordenada
func (l *Lista) insereOrd(ch int) {
	ant, no = l.busca(ch)

	if no != nil {
		return
	}

	novoNo = &No{chave: ch}
	if ant == nil {
		novoNo.prox = l.cab
		l.cab = novoNo
	} else {
		novoNo.prox = ant.prox
		ant.prox = novoNo
	}
}

// remoção de uma lista ordenada
func (l *Lista) remove(ch int) *No {
	ant, no = l.busca(ch)

	if no == nil {
		return nil
	}

	if ant == nil {
		l.cab = no.prox
	} else {
		ant.prox = no.prox
	}

	return no
}

func main() {
	var l Lista

	// fmt.Println(l.busca(2))
	l.insereOrd(2)
	l.insereOrd(4)
	l.insereOrd(6)
	l.insereOrd(3)
	l.insereOrd(1)
	// l.exibe()
	fmt.Println(l.remove(5))
	fmt.Println(l.remove(3))
	fmt.Println(l.remove(6))
	fmt.Println(l.remove(1))
	fmt.Println("-------------")
	l.exibe()
	// fmt.Println(l.busca(2))
	// fmt.Println(l.busca(1))
	// fmt.Println(l.busca(8))
	// fmt.Println(l.busca(5))
	// fmt.Println(l.busca(4))
	// fmt.Println(l.buscaSimples(2))
	// fmt.Println(l.buscaSimples(4))
	// fmt.Println(l.buscaSimples(6))
}

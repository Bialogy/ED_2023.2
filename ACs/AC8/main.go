package main

import "fmt"

// Ana Beatriz Lessa && João Pedro Alencar
// Árvore binária de busca, exemplo de um dicionário
// fonte: https://copyprogramming.com/howto/binary-search-tree-implementation-in-golang
// Para buscar outros valores, alterar a linha 78 com a chave que deseja buscar.

type No struct {
	Chave int
	Valor string
	Esq   *No
	Dir   *No
}

// Buscar um valor no dicionário
func Busca(raiz *No, chave int) string {
	if raiz == nil {
		return "Chave não encontrada"
	}

	if chave == raiz.Chave {
		return raiz.Valor
	} else if chave < raiz.Chave {
		if raiz.Esq == nil {
			return "Chave não encontrada"
		}
		return Busca(raiz.Esq, chave)
	} else {
		if raiz.Dir == nil {
			return "Chave não encontrada"
		}
		return Busca(raiz.Dir, chave)
	}
}

// Insere uma chave e seu valor no dicionário
func Inserir(raiz *No, chave int, valor string) *No {
	if raiz == nil {
		return &No{Chave: chave, Valor: valor, Esq: nil, Dir: nil}
	}

	resultado := Busca(raiz, chave)

	if resultado != "Chave não encontrada" {
		fmt.Println("Inválido: já existe no dicionário")
		return raiz
	}

	novoNo := &No{Chave: chave, Valor: valor, Esq: nil, Dir: nil}

	if chave < raiz.Chave {
		if raiz.Esq == nil {
			raiz.Esq = novoNo
		} else {
			Inserir(raiz.Esq, chave, valor)
		}
	} else {
		if raiz.Dir == nil {
			raiz.Dir = novoNo
		} else {
			Inserir(raiz.Dir, chave, valor)
		}
	}

	return raiz
}

func main() {

	dicionario := &No{Chave: 50, Valor: "maçã", Esq: nil, Dir: nil}

	Inserir(dicionario, 30, "banana")
	Inserir(dicionario, 20, "laranja")
	Inserir(dicionario, 10, "pêra")
	Inserir(dicionario, 40, "uva")

	chaveBusca := 40
	resultadoBusca := Busca(dicionario, chaveBusca)

	if resultadoBusca != "Chave não encontrada" {
		fmt.Printf("Valor da chave %d: %s\n", chaveBusca, resultadoBusca)
	} else {
		fmt.Printf("Chave %d não encontrada no dicionário.\n", chaveBusca)
	}

}

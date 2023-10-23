package produtos

var TotalProdutos = 0

type Produto struct {
	Id int
	Nome string
}

func (p *Produto) DefinirId() {
	TotalProdutos++
	p.Id = TotalProdutos
}

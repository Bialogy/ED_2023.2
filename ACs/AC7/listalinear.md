# Listas Lineares

Uma lista linear é uma estrutura de dados em que seus elementos se organizam de forma linear. Cada elemento aponta para o próximo elemento até o final da lista.

No caso de uma lista linear duplamnte encadeada, cada elemento aponta para o próximo assim como para o seu anterior. Logo, é possível percorrer tal lista em ambas as direções.

### Exibição dos nós em uma lista linear duplamente encadeada:

##### Lista Linear Simplesmente Encadeada:

Nó1 ( **5** ) ---> Nó2 ( **8** ) ---> Nó3 ( **13** ) ---> Nó4 ( **21** ) ---> Nó5 ( **34** )

##### Lista Linear Duplamente Encadeada.

Nó1 ( **5** ) <---> Nó2 ( **8** ) <---> Nó3 ( **13** ) <---> Nó4 ( **21** ) <---> Nó5 ( **34** )

### Busca de um nó:

Para buscar um nó com um valor específico, é necessário percorrer a lista linear sabendo que ela está ordenada. Caso o primeiro elemento tenha um valor maior do que o valor de busca, não existe este elemento na lista. Caso contrário, ele deve passar para o próximo elemento até que a busca encerre.

```go
função BuscaNó(lista, valorBusca)

    NóAtual = lista.primeiro

    Enquanto NóAtual != 0:
        Se NóAtual.Nó = valorBusca:
            Retorne NóAtual
        Senão Se NóAtual.Nó > valorBusca:
            Retorne Nil
        Senão:
            NóAtual = NóAtual.próximo
    
    Retorne Nil

```

### Inserção de um nó

No caso de uma inserção de um nó, é necessário primeiramente fazer uma busca pela lista ordenada para saber onde o elemento deve encaixar.

```go
função InserirNó(lista, valorInserção)

    NovoNó = NovoNó
    NovoNó.Nó = valorInserção

    Se lista.estáVazia:
        lista.primeiro = NovoNó
        lista.ultimo = NovoNó
    Senão Se valorInserção < lista.primeiro.Nó:
        NovoNó.próximo = lista.primeiro
        lista.primeiro.anterior = NovoNó
        lista.primeiro = NovoNó
    Senão
        NóAtual = lista.primeiro
        Enquanto NóAtual.Próximo != Nil E NóAtual.Próximo.Nó < valoInserção:
            NóAtual = NóAtual.próximo
        
        NovoNó.próximo = NóAtual.Próximo
        Se NóAtual.próximo != Nil:
            NóAtual.próximo.anterior = NovoNó
        Senão:
            lista.último = NovoNó

        NovoNó.anterior = NóAtual
        NóAtual.próximo = NovoNó
```


### Remoção de um nó

Para a remoção de um nó é preciso buscar na lista pelo seu valor e atualizar os ponteiros anteriores e posteriores para apontarem para si mesmos. Se o valor a ser removido estiver no início ou fim da lista, seus anteriores ou posteriores devem ser atualizados para receberem um valor nulo respectivamente.

```go
função InserirNó(lista, valorRemoção)

    Se lista.estáVazia:
        retornar
    Senão:
        NóAtual = lista.primeiro
        Enquanto NóAtual != Nil E NóAtual.Nó < valorRemoção:
            NóAtual = NóAtual.próximo
        Se NóAtual != Nil E NóAtual.Nó == valorRemoção:
            Se NóAtual == lista.primeiro:
                lista.primeiro = NóAtual.próximo
                se lista.primeiro != Nil :
                lista.primeiro.anterior = Nil
            Senão Se NóAtual == lista.ultimo:
                lista.ultimo = NóAtual.anterior
                lista.ultimo.próximo = Nil
            Senão:
                NóAtual.anterior.próximo = NóAtual.próximo
                NóAtual.próximo.anterior = NóAtual.anterior
```
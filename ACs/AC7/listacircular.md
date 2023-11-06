# Listas Circulares

Listas Circulares são exemplos de listas em que cada elemento está ligado ao próximo e o último elemento está ligado ao primeiro. Logo, a lista cria uma corrente, formando um ciclo. Dentre as listas circulares, há dois tipos: simples e dupla.

Uma lista circular simples permite percorrê-la, em um ciclo, em apenas uma direção. Seus nós possuem ponteiros apenas para o elemento seguinte. Enquanto isso, uma lista circular dupla permite que se percorra em ambas as direções, possuindo ponteiros que apontando para elementos anteriores e posteriores.

### Exibição dos nós em uma lista circular:

##### Lista Circular Simples

"1" ---> "2" ---> "3" ---> "4" ---> "5" ---> "1" ---> "2" ---> "3" ---> ...

```go
Struct do Nó:
    Nó: Inteiro
    ProxNó: Ponteiro de Nó

Nó1 = Outro Nó
Nó1.Nó = 1

Nó = Outro Nó
Nó2.Nó = 2

Nó = Outro Nó
Nó3.Nó = 3

Nó = Outro Nó
Nó4.Nó = 4

Nó = Outro Nó
Nó5.Nó = 5

Nó1.ProxNó = Nó2
Nó2.ProxNó = Nó3
Nó3.ProxNó = Nó4
Nó4.ProxNó = Nó5
Nó5.ProxNó = Nó1

AtualNó = Nó1
For i de 1 até 10:
    Print(AtualNó.Nó)
    AtualNó = AtualNó.ProxNó
Break
```

##### Lista Circular Dupla

"1" <---> "2" <---> "3" <---> "4" <---> "5" <---> "1" <---> "2" <---> "3" <---> ...

```go
Struct do Nó:
    Nó: Inteiro
    ProxNó: Ponteiro de Nó
    NóAnt: Ponteiro de Nó

Nó1 = Outro Nó
Nó1.Nó = 1

Nó = Outro Nó
Nó2.Nó = 2

Nó = Outro Nó
Nó3.Nó = 3

Nó = Outro Nó
Nó4.Nó = 4

Nó = Outro Nó
Nó5.Nó = 5

Nó1.ProxNó = Nó2
Nó1.NóAnt = Nó5
Nó2.ProxNó = Nó3
Nó2.NóAnt = Nó1
Nó3.ProxNó = Nó4
Nó3.NóAnt = Nó2
Nó4.ProxNó = Nó5
Nó4.NóAnt = Nó3
Nó5.ProxNó = Nó1
Nó5.NóAnt = Nó4

AtualNó = Nó1
For i de 1 até 10:
    Print(AtualNó.Nó)
    AtualNó = AtualNó.ProxNó
Break
```

### Inserção de um nó no início da lista circular simplesmente encadeada

Ao inserir um novo elemento no início de uma lista circular, o mesmo deve ser inserido com um ponteiro para o elemento que anteriormente ocupava o primeiro lugar. O último elemento também deve ser atualizado para apontar para o novo elemento.

```go
função InserirNoInicio(lista, nó):

    NovoNó.ProxNó = NovoNó
    NovoNó.Nó = Nó

    NovoNó.ProxNó = lista.Primeiro
    lista.ultimo.próximo = NovoNó
    lista.primeiro = NovoNó

    retorna lista
```

### Exclusão do primeiro nó da lista circular simplesmente encadeada

Ao excluir o primeiro elemento de uma lista circular, é preciso atualizar o antigo segundo elemento para que se torne o primeiro e também é necessário fazer com que o último elemento aponte agora para o novo primeiro. 

```go
função ExcluirNoInicio(lista):
    
    PrimeiroNó = lista.primeiro
    lista.primeiro = lista.primeiro.próximo 
    lista.ultimo.próximo = lista.primeiro

    retorna lista
```



Fonte(s)
[Programiz](https://www.programiz.com/dsa/circular-linked-list)
[Logicmojo](https://logicmojo.com/circular-linked-list-problem#memory)
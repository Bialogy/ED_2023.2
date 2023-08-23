// 1- Organize o projeto em pacotes: um pacote para o tipo Contato,
// um para a operação sobre os arrays e um arquivo main.go.
// 2- Refatore as funções que retornam arrays para que elas utilizem ponteiros.
// 3- Crie uma nova função editaEmail, que recebe o índice do elemento
// no array e altera o e-mail do elemento indicado.
// 4- Atualize a interface por linha de comando
// para incluir a opção de exibir todos os contatos salvos.
// Cada contato deve ser exibido em uma linha,
// precedido pelo número do índice daquele elemento.
// 5- Atualize a interface por linha de comando para poder editar o e-mail
// de um contato previamente saldo.
// A interface deve exibir os contatos armazenados
// e pedir para o usuário indicar o índice do contato que quer ser alterado.
// Em seguida, o programa pede o novo e-mail e chama a função editaEmail,
// implementada no exercício 3.

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Olá, mundo!")

	c := ctt.Contato{Nome: "", Email: ""}

}
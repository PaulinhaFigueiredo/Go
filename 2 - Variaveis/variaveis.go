package main

import "fmt"

func main() {

	//Modo de declaraçao de variavel em GO
	var variavel1 string = "variavel1"

	//Modo de declaraçao utilizado em GO por convençao
	variavel2 := "variavel 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)
	fmt.Println()

	//Declaraçao de varias variaveis
	var (
		variavel3 string = "variavel3"
		variavel4 string = "variavel4"
		variavel5 string = "variavel5"
	)

	fmt.Println(variavel3, variavel4, variavel5)
	fmt.Println()

	//Declaraçao de varias variaveis com interferencia de tipos
	variavel6, variavel7 := "variavel 6", "variavel 7"
	fmt.Println(variavel6, variavel7)

	/*Constante em GO , segue as mesmas regras das variaveis
	porem nao altera o valor da costante*/
	const constante1 string = "Constante 1"
	fmt.Println(constante1)

}

/*
O Go trabalha com tipo de variavel implicita, se voce passa dentro da variavel
um valor string por exemplo, ele entende que sua variavel vai receber o tipo string.

EX1:  nomePessoa := "Raquel"

No EX1 Não precisamos declarar o tipo que a variavel nome pessoa é,
pq passamos um valor do tipo string e o Go ja entendeu que se trata
de uma variavel String

EX2: valorSaldo := 100.23

No EX2 estamos passando um valor flutuante, e o Go entende que
estamos trabalhando com uma variavel do tipo Float


*/

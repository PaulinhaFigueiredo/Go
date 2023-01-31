package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("devbook@gmail.com")
	fmt.Println(erro)
}

/*
Para utilizar algo importado, exemplo o github.com na linha 7
usamos a palavra depois da ultima barra nesse caso
utilizamos o checkmail, linha 14
*/

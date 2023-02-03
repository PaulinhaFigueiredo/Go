package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int64 = 1000
	fmt.Println(numero)

	var numero2 int32 = 2000
	fmt.Println(numero2)

	//alias
	//INT32 = RUNE
	var numero3 rune = 3000
	fmt.Println(numero3)

	//BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)
}

/*
	Existem 4 tipos de numeros inteiros em GO
	int8, int16, int32, int64
	a diferença é so o tamanho

	Existe tambem o: int
	Ele sozinho quando vc nao especifica, ele usa a arquitetura
	do computador como base.
	Se o seu computador for windows64 ele vai criar um int64,
	ou for windows32 ele cria um int32

*/

func main2() {
	var numero1 float32 = 123.45
	fmt.Println(numero1)

	var numero2 float64 = 2000.20
	fmt.Println(numero2)

	numeroReal := 12345.68
	fmt.Println(numeroReal)
}

/*
	Inferencia de tipos no float funciona da mesma maneira que no INT
	Ele pega a arquitetura do seu computador e lança, se for windows64 ele lança float64
	EX: linha 42

	No Go sempre usamos aspas duplas " ".
	O go se vc iniciar uma variavel vazia que seja numero, ele inicializa
	ele com zero, se for string ela inicia como null.

*/

func main3() {

	var booleano bool
	fmt.Println(booleano)

}
/*
	No booleano Se for declarar uma variavel como falsa nao precisa 
	passar valor, apenas declarar ela como na linha 59
*/


func main4() {

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}

/*

	Criando um erro em GO. 

	linha 71: erro em azul claro é o nome da variavel, error com r no final(verde) é o tipo da variavel.
	erros com S no final é um pacote do Go (importado na linha 4) dando um erros.New() é possivel passar a msg de erro
	linha 75

*/
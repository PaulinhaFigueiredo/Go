package auxiliar

import "fmt"

//Escrever registra uma msg na tela
func Escrever() {
	fmt.Println("Escrevendo do pacote auxiliar")
	escrever2()
}

/*
Se uma funcao começa com letra maiuscula significa que ela é
uma funçao publica. Ela pode ser exportada por outros pacotes.
*/

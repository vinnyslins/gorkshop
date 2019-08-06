package main

import "fmt"

func main() {
	// o array é uma estrutura pouco utilizada
	var arrayInicial [5]int
	arrayInicial[0] = 1

	fmt.Println(arrayInicial)
	fmt.Println(len(arrayInicial))

	// é possível criar um array atribuíndo seu
	// tamanho dinamicamente
	var arrayLimiteNaHora = [...]int{1, 2, 3, 4, 5}
	fmt.Println(arrayLimiteNaHora)
}

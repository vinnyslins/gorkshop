package main

import "fmt"

func main() {
	// não informando o tamanho, teremos um slice
	var sliceInicial []int
	fmt.Println(sliceInicial)
	sliceInicial = append(sliceInicial, 1, 2, 3, 4)
	fmt.Println(sliceInicial)

	// para fatiar o slice
	fmt.Println(sliceInicial[:2])

	// make(
	// 	slice inicial
	// 	tamanho do slice inicial,
	// 	máximo do tamanho (exponencial)
	// )
	novoSlice := make([]int, 5, 10)
	println(novoSlice)
}

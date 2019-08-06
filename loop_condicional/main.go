package main

import "fmt"

func main() {
	i := 0

	for i < 10 {
		fmt.Println(i)
		i++
	}

	for x := 0; x < 10; x++ {
		fmt.Println(x)
	}

	// CUIDADO
	// um loop infinito pode estourar memória
	// uma vez que o go vai disparar múltiplas
	// threads para executar o loop
	a := 0
	for {
		a++
		if a == 10 {
			break
		}
	}

	// para percorrer um slice
	meuSlice := []int{1, 2, 3, 4, 5}
	for _, value := range meuSlice {
		fmt.Println(value)
	}
}

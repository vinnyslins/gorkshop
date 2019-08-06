package main

import "fmt"

func main() {
	// o map é uma lista de chave/valor
	// primeiroMap := map[string]int{}
	primeiroMap := map[string]int{
		"olá":   1,
		"mundo": 2,
	}

	primeiroMap["novo"] = 3
	fmt.Println(primeiroMap["novo"])

	novoMap := make(map[string]int)
	novoMap["teste"] = 1
	fmt.Println(novoMap["teste"])

	mapAninhado := map[string]map[int]bool{
		"oi": map[int]bool{
			1: true,
		},
		"mundo": {
			2: false,
		},
	}
	fmt.Println(mapAninhado["oi"][1])
}

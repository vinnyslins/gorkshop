package main

import "aplicacaomain/model"

func main() {
	personagem := model.Personagem{}
	personagem.SetName("Vinnys")
	personagem.Falar()
}

package main

type Personagem struct {
	Name     string
	Age      int
	Gender   string
	Endereco Endereco
}

type Endereco struct {
	Logradoro string
	numero    int
}

func main() {
	personagem := Personagem{
		Name:   "Vinnys",
		Age:    20,
		Gender: "Masculno",
	}

	personagem.Name = "Francisco"
	personagem.Endereco.Logradoro = "Rua dois"
	personagem.Endereco.numero = 2
}

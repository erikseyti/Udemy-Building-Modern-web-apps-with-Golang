package main

import "log"

func main () {
	animals := make(map[string]string)
	animals["dog"] ="Fido"
	animals["cat"] = "Fluffly"

	for animalType, animal :=range animals {
		log.Println(animal, " is a(n) ", animalType)

	}

	var firsLine = "Once upon a midnight dreary"
	// Ao voce alterar o valor de uma variavel em Go.
	// O objeto anterior na verdade ele eh destruido e em seguida.
	// criasse um novo com o novo valor
	firsLine = "x"

	// nesse caso ele nao retorna diretamente a letra, mas sim o valor em bytes
	// da letra correspondente
	for i, l := range firsLine{
		log.Println(i, ":", l)
	}
}
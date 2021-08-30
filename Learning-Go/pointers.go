package main

import "log"

func main() {
	var myString string
	myString = "Green"

	log.Println("My string is set to ", myString)
	changeUsingPointer(&myString)
	log.Println("after func call My string is set to ", myString)
}

func changeUsingPointer(s *string) {
	// essa funcao muda o valor de uma variavel a partir da localizacao
	// do seu ponteiro na memoria
	log.Println("s is set to", s)
	newValue := "Red"
	*s = newValue
}

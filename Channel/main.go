package main

import (
	"log"

	"github.com/erikseyti/myniceprogram2/helpers"
)

const numPool = 1000

func CalculateValue(intChan chan int) {
	randomNumbers := helpers.RandomNumber(numPool)
	intChan <- randomNumbers

}

func main() {
	intChan :=make(chan int)
	// apos conseguir pegar o valor para intChan, imediamente com comando defer
	// Fecha a sua respectiva conexao
	defer close(intChan)

	// roda ele como uma rotina Go, onde ele executa simultaneamente todas as rotinas
	// do tipo go.
	go CalculateValue(intChan)

	num := <- intChan
	log.Println(num)

}

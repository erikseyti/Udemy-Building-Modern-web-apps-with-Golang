package main

import "log"

func main() {
	for i:= 0; i<=10; i++ {
		log.Println(i)
	}

	animals := []string {"dog", "fish", "horse", "cow", "cat"}

	for i, animal := range animals {
		log.Println(i, animal)

	}

	// implementado underline para indicar que o valor do index n sera utilizado
	// dentro desse for
	for _, animal := range animals {
		log.Println( animal)

	}
}
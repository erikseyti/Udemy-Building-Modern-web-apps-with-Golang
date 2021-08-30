package main

import "log"

func main() {

	myVar := "cat"
	// No Go, o switch case somente roda ate achar a primeira istancia que ele encontra.
	// Por isso nao eh necessario implementar o uso do break apos cada case dentro do switch
	switch myVar {
	case "cat":
		log.Println("cat is set to cat")
	case "dog":
		log.Println("cat is set to dog")
	case "fish":
		log.Println("cat is set to fish")
	default:
		log.Println("cat is something else")
	}

}

package main

import "log"

func main() {
	// Tipo booleano em Go
	// outra forma de declarar:
	// var isTrue bool
	// isTrue = true
	isTrue := true

	// or isTrue == true
	if isTrue {
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}

	cat := "cat"

	if cat == "cat2" {
		log.Println("Cat is cat")
	} else {
		log.Println("Cat is not cat")
	}

	myNum := 100
	isTrue2 := false

	if myNum > 99 && !isTrue2 {
		log.Println("myNum is Greater than 99 and isTrue is set to true")
	} else if myNum < 100 && isTrue2 {
		log.Println("1")
	} else if myNum == 101 || isTrue2 {
		log.Println("2")
	} else if myNum > 1000 && isTrue2 == false {
		log.Println("3")
	}

}

package main

import (
	"log"
	"time"
)

// var s string = "seven"
// var v = "eight"

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	// Tipo dentro do Go para registrar tempo, podendo ser literalmente um valor
	// em tempo, data ou dia
	BirthDate time.Time
}

func main() {
	// usamos letra maiuscula para indicar que a varivel é publica e pode ser
	// utilizada fora dessa arquivo.
	// caso utilize em letra minuscula, ela é privada e somente pode ser usada
	// no escopo desse arquivo
	user := User {
		FirstName: "Trevor",
		LastName: "Scoot",
		PhoneNumber: "1 555 555-1212",
	}

	log.Println(user.FirstName, user.LastName, user.BirthDate, user.PhoneNumber)
	// var s2 = "six"

	// log.Println("s is: ", s)
	// log.Println("s2 is: ", s2)

	// saySomething("xxx")
}

// func saySomething(s3 string) (string, string) {
// 	log.Println("Say Something s is: ", s)
// 	return s3, "World"
// }

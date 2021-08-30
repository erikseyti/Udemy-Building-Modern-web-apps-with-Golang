package main

import "log"

func main() {
	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}
	var users []User
	users = append(users, User{"John", "Smith", "John@smith.com", 30})
	users = append(users, User{"Melody", "Jones", "Melody@jones.com", 20})
	users = append(users, User{"Annita", "Vanessa", "annitta@vanessa.com", 25})
	users = append(users, User{"Peter", "Parker", "IamNot@Spiderman.web", 28})

	for _, l := range users {
		log.Println(l.FirstName, l.LastName,l.Email, l.Age)

	}
}

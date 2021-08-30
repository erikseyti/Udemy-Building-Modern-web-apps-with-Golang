package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	// Maps nao mantem a ordem, por decisao do time que fez o Go.
	myMap := make(map[string]string)

	myMap["dog"] = "Samson"
	myMap["other-dog"] = "Cassie"

	myMap["dog"] = "fido"

	log.Println(myMap["dog"])
	log.Println(myMap["other-dog"])

	myMap2 := make(map[string]int)

	myMap2["First"] = 1
	myMap2["Second"] = 2

	log.Println(myMap2["First"])
	log.Println(myMap2["Second"])

	myMap3 := make(map[string]User)

	me := User{
		FirstName: "Trevor",
		LastName:  "Scoot",
	}

	myMap3["me"] = me
	log.Println(myMap3["me"])
	log.Println(myMap3["me"].FirstName)

	var myNewVar float32
	myNewVar = 11.1

	log.Println(myNewVar)

	



	// var myString string
	// var myInt int

	// myString = "Hi!"
	// myInt = 11

	// mySecondString := "another string"

	// log.Println("All my variables: ", myString, myInt, mySecondString)

}

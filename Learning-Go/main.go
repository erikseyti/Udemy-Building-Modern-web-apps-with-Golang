package main

import "fmt"

func main() {
	fmt.Println("Hello, World.")

	var whatToSay string
	var i int

	whatToSay = "Goodbye, Creul World"

	fmt.Println(whatToSay)

	i = 7

	fmt.Println("i is set to", i)

	whatWasSaid, theOtherThingThatWasSaid := saySomething()

	fmt.Println("The Function returned", whatWasSaid, theOtherThingThatWasSaid)
}

func saySomething() (string, string) {
	return "something", "else"
}
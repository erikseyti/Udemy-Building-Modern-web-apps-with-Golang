package main

import (
	"log"

	"github.com/erikseyti/myniceprogram/helpers"
)

func main () {
	log.Println("Hello")

	var myVar helpers.SomeType
	myVar.TypeName = "Some name"

	log.Println(myVar.TypeName)
}


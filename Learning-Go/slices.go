package main

import (
	"log"
	"sort"
)

func main() {
	// Maneira de manejar uma serie de dados em GO, 
	// Similar a listas em Python e Javascript
	var mySlice []string

	mySlice = append(mySlice, "Trevor")
	mySlice = append(mySlice, "John")
	mySlice = append(mySlice, "Mary")

	var mySlice2 []int

	mySlice2 = append(mySlice2, 1)
	mySlice2 = append(mySlice2, 8)
	mySlice2 = append(mySlice2, 2)

	log.Println(mySlice)
	log.Println(mySlice2)

	sort.Ints(mySlice2)
	log.Println(mySlice2)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	log.Println(numbers)

	log.Println(numbers[0:2])
	log.Println(numbers[6:9])
	log.Println(numbers[5:])
	log.Println(numbers[:7])

	names := []string{"one", "seven", "fish", "cat"}
	log.Println(names)
}

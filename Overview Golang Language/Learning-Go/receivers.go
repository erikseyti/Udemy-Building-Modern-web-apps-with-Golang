package main

import "log"

type myStruct struct {
	FirstName string
}

// Esse tipo de anotacao eh chamado de receiver,
// Seria uma funcao diretamente associada a um struct
// o paralelo a orientacao a objeto seria um metodo
// OBS: Go n utiliza o conceito de orientacao a objeto na linguagem.
func (m *myStruct) printFirstName() string {
	return m.FirstName + "!"
}

func main() {
	var myVar myStruct
	myVar.FirstName = "John"

	myVar2:= myStruct{
		FirstName: "Mary",
	}

	// log.Println("MyVar is set to", myVar.FirstName)
	// log.Println("MyVar2 is set to", myVar2.FirstName)
	log.Println("MyVar is set to", myVar.printFirstName())
	log.Println("MyVar2 is set to", myVar2.printFirstName())

}
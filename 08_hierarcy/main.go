package main

import "fmt"

func main() {
	human := Human{"Dario", "Lopez", 24}
	fmt.Printf("%v \n", human.GetName())
	student := Student{Human: Human{"Marcelo", "Perez", 24}, Legajo: "123123"}
	fmt.Printf("%v \n", student.GetName())
	teacher := Teacher{Human: Human{"Ricardo", "Fort", 24}, Materia: "Matematicas"}
	fmt.Printf("%v \n", teacher.Human.GetName())
}

type Human struct {
	FirstName string
	LastName  string
	Age       int32
}

type Student struct {
	Human
	Legajo string
}

type Teacher struct {
	Human
	Materia string
}

func (human *Human) GetName() string {
	return human.FirstName + " " + human.LastName
}

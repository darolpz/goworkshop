package main

import "fmt"

func main() {

	var human Human
	student := Student{"Daro", "Lopez", 24}
	teacher := Teacher{"Sergio", "Marquina", 32, "Historia"}
	/* Student y teacher puede asignarse a variables de tipo Human */
	human = student
	human = teacher

	fmt.Printf("%v \n", human.GetAge())

	/* Tipo Student es valido como tipo Human */
	sayName(student)

}

func sayName(human Human) {
	fmt.Printf("Mi nombre es %v \n", human.GetName())
}

type Human interface {
	GetName() string
	GetAge() int32
}

type Student struct {
	FirstName string
	LastName  string
	Age       int32
}

type Teacher struct {
	FirstName string
	LastName  string
	Age       int32
	Subject   string
}

func (human Student) GetName() string {
	return human.FirstName + human.LastName
}

func (human Teacher) GetName() string {
	return human.FirstName + human.LastName
}

func (human Student) GetAge() int32 {
	return human.Age
}

func (human Teacher) GetAge() int32 {
	return human.Age
}

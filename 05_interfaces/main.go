package main

import "fmt"

func main() {

	var human Human
	student := Student{"Daro", "Lopez", 24}
	teacher := Teacher{"Sergio", "Marquina", 32, "Historia"}

	human = student
	sayName(human)
	human = teacher
	sayName(human)
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

func (student Student) GetName() string {
	return "Alumno " + student.FirstName + " " + student.LastName
}

func (teacher Teacher) GetName() string {
	return "Profesor " + teacher.FirstName + " " + teacher.LastName
}

func (student Student) GetAge() int32 {
	return student.Age
}

func (teacher Teacher) GetAge() int32 {
	return teacher.Age
}

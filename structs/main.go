package main

import "github.com/darolpz/workshop/structs/person"

func main() {

	person := person.Person{FirstName: "Dario", Lastname: "Lopez"}

	person.Saludar()
}

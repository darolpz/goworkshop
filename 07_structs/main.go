package main

import "github.com/darolpz/workshop/07_structs/person"

func main() {

	person := person.Person{FirstName: "Dario", Lastname: "Lopez"}

	person.Saludar()
}

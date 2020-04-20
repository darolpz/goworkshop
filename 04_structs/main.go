package main

import "github.com/darolpz/workshop/04_structs/person"

func main() {

	person := person.Person{FirstName: "Dario", Lastname: "Lopez"}

	person.Saludar()
}

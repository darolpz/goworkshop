package person

import "fmt"

type Person struct {
	FirstName string
	Lastname  string
}

func (person *Person) Saludar() {
	fmt.Printf("Mi nombre es %s %s", person.FirstName, person.Lastname)
}

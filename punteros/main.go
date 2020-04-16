package main

import "fmt"

func main() {
	/*
		Se declaran con & y se acceden con *
	*/
	number := 5
	pointer := &number
	fmt.Printf("Direccion de memoria %v\n", *pointer)
	fmt.Printf("Valor %v\n", *pointer)
	number = -114
	fmt.Printf("Valor %v\n", *pointer)
	*pointer = 25
	fmt.Printf("Valor %v\n", *pointer)
	fmt.Println("")

	user := User{"Dario", 24}
	fmt.Printf("%v\n", user)
	user.changeAgePointer(25)
	fmt.Printf("%v\n", user)

	user.changeAgeNoPointer(30)
	fmt.Printf("%v\n", user)
}

/* Usar punteros evitar la copia del valor receptor en las llamadas al método (más eficiente) */
/* El método puede modificar el valor del receptor.  */
/* Si es un metodo solo de lectura no importa si se pasa una referencia o copia */
type User struct {
	Name string
	Age  int32
}

func (user *User) changeAgePointer(newAge int32) {
	user.Age = newAge
}

/* En vez de pasarse una referencia a la variable user se pasa una copia que no afectara al programa principal */
func (user User) changeAgeNoPointer(newAge int32) {
	user.Age = newAge
}

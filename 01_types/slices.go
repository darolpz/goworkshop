package main

import "fmt"

func main() {
	//Slices
	s := []int{1, 2, 3, 4}
	fmt.Printf("%v\n", s)
	s = append(s, 5)
	fmt.Printf("%v\n", s)
	/* Tipo de dato, longitud inicial y capacidad */
	/* Su valor 0 es nil */
	slice := make([]int, 1, 1)
	fmt.Printf("%v\n", slice)
	/* No tiene problemas con agregar elementos mas alla de su capacidad, pero necesita una reasignacion de memoria */
	slice = append(slice, 5)
	fmt.Printf("%v\n", slice)
}

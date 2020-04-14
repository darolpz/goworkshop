package main

import (
	"fmt"

	"github.com/darolpz/workshop/types/helper"
)

func main() {
	//Tipos de datos en Golang
	/*
	   bool

	   string

	   int  int8  int16  int32  int64
	   uint uint8 uint16 uint32 uint64 uintptr

	   byte // alias for uint8

	   rune // alias for int32
	        // represents a Unicode code point

	   float32 float64

	   complex64 complex128
	*/

	// Formas de declarar variables
	var string1 string
	string1 = "String 1"
	string2 := "String 2"
	var string3 string = "String 3"
	fmt.Printf("%v %v %v \n", string1, string2, string3)

	// Arreglos
	var arreglo [5]int
	arreglo[0] = 1
	fmt.Printf("%v\n", arreglo)

	//Valor 0
	var val0 string
	fmt.Printf("%v\n", val0 == "")

	// Constantes
	const pi = 3.14
	fmt.Printf("%v\n", pi)
	// pi = 5

	//If
	if boolean := false; boolean {
		fmt.Printf("%v", boolean)
	}

	//Funciones con dos resultados y modificadores de accesos
	sin, cos := helper.SinCosHelper(0)
	fmt.Printf("%v  %v \n ", sin, cos)

	//Map
	m := map[string]Coords{
		"Bell Labs": {"40.68433", "-74.39967"},
		"Google":    {"37.42202", "-122.08408"},
		"Argentina": {"-34", "-58"}, //Si se ordena los structs de esta forma debe haber una coma al final
	}
	mapVar := make(map[string]Coords)
	mapVar["P0"] = Coords{"0", "0"}
	mapVar["Argentina"] = Coords{"-34", "-58"}
	fmt.Printf("%v %v\n", m, mapVar)

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

type Coords struct {
	Latitude  string
	Longitude string
}

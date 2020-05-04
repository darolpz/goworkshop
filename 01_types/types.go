package main

import (
	"fmt"

	"github.com/darolpz/workshop/01_types/helper"
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
	fmt.Printf("%T %T %T\n", string1, string2, string3)

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

}

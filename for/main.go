package main

import "fmt"

func main() {

	//normal
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Printf("\n")

	// sin asignacion
	index0 := 0
	for ; index0 < 5; index0++ {
		fmt.Println(index0)
	}

	fmt.Printf("\n")
	//While
	index := 0
	for index < 5 {
		fmt.Println(index)
		index++
	}

	// Bucle infinito
	/* for{

	} */
	fmt.Printf("\n")
	arreglo := [5]int{1, 2, 3, 4, 5}

	for index := range arreglo {
		fmt.Println(index)
	}
	fmt.Printf("\n")
	for _, value := range arreglo {
		fmt.Println(value)
	}

	fmt.Println(arreglo[0])
}

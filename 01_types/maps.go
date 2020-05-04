package main

import "fmt"

func main(){
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
}
type Coords struct {
	Latitude  string
	Longitude string
}

package hello

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	fmt.Printf("Validating")
	name := "Daro"
	want := "Hello Daro"
	if result := Hello(name); result != want {
		t.Errorf("Hello() = %q, want %q", result, want)
	}
}

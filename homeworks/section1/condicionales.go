package main

import (
	"fmt"
)

func main() {

	license := true
	age := 25

	if age > 15 && license {
		fmt.Println("Puede seguir avanzando")
	}

	if age <= 15 && !license {
		fmt.Println("No puede seguir circulando")
	}

	switch {
	case license && age >= 19:
		fmt.Println("Pruede seguir avanzando")
	case !license && age >= 19:
		fmt.Println("No puede seguir circulando")
	case license && age <= 15:
		fmt.Println("No puede seguir circulando")
	}
}

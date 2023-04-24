package main

import (
	"fmt"
)

func main() {

	var array []string
	fmt.Println("Seleccione valores, se sale con cero")

	for {
		var value string
		fmt.Scanf("%s", &value)

		if value == "0" {
			break
		}

		array = append(array, value)
	}

	fmt.Println("Los valores del array son: ", array)
}

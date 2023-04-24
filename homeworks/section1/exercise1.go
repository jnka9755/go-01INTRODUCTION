package main

import (
	"fmt"
)

func main() {

	array := [5]int{4, 2, 5, 6, 7}

	for i := range array {

		array[i] += 20
	}

	fmt.Println("Los valores del array son: ", array)
}

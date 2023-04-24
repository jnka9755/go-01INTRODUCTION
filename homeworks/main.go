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

		if value != "" {
			switch value {
			case "10":
				array = append(array, "notebook")
			case "15":
				array = append(array, "tv")
			case "21":
				array = append(array, "heladera")
			case "27":
				array = append(array, "monitor")
			case "35":
				array = append(array, "camara")
			default:
				array = append(array, "notFound")
			}
		}
	}

	fmt.Println("Los valores del array son: ", array)
}

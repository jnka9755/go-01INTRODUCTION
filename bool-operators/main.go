package main

import (
	"fmt"
)

func main() {

	age := 32

	fmt.Println("Operators")

	fmt.Println(age > 30)
	fmt.Println(age < 30)
	fmt.Println(age <= 30)
	fmt.Println(age >= 30)
	fmt.Println(age == 30)

	fmt.Println()

	// OR
	fmt.Println(age < 32 || age == 32)
	fmt.Println(age < 32 || age == 33)
	fmt.Println(age < 40 || age == 33)

	fmt.Println()

	//AND
	fmt.Println(age < 32 && age == 32)
	fmt.Println(age < 32 && age == 33)
	fmt.Println(age < 40 && age == 32)

	fmt.Println()

	//NOT
	fmt.Println(!(age > 30))
	fmt.Println(!(age < 30))
	fmt.Println(!(age <= 30))
	fmt.Println(!(age >= 30))
	fmt.Println(!(age == 30))

	fmt.Println()

	fmt.Println(age < 25 && age == 32 || age < 40)
	fmt.Println(age < 25 && (age == 32 || age < 40))

	age = 20

	if age > 18 {
		fmt.Printf("%d is higer than 18\n", age)
	}

	boolVar := true

	if boolVar {
		fmt.Println("is true")
	} else {
		fmt.Println("is false")
	}

	if value := true; value {
		fmt.Println("is true")
	}

	number := 3

	switch number {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("undefined number")
	}

	switch {
	case number > 0:
		fmt.Println("positive")
	case number < 0:
		fmt.Println("negative")
	case number == 0:
		fmt.Println("zero")
	}
}

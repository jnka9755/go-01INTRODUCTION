package main

import (
	"fmt"

	"github.com/jnka9755/go-01INTRODUCTION/functions/function"
)

func main() {
	fmt.Println(function.Add(3, 4))

	function.RepeatString(5, "Juan")

	fmt.Println()

	value, error := function.Calc(function.SUM, 5, 3)

	if error != nil {
		fmt.Println("Error: ", error)
	} else {
		fmt.Println("Value: ", value)
	}

	value, error = function.Calc(function.DIV, 10, 0)

	if error != nil {
		fmt.Println("Error: ", error)
	} else {
		fmt.Println("Value: ", value)
	}

	fmt.Println()
	x, y := function.Split(2)
	fmt.Println(x, y)

	fmt.Println()
	value = float64(function.MSum(1, 2, 3, 4, 5, 6, 7, 8, 9))
	fmt.Println(value)

	fmt.Println()

	value, error = function.MOperations(function.SUM, 1, 2, 3)
	fmt.Println(value, error)

	value, error = function.MOperations(function.SUB, 1, 2, 3)
	fmt.Println(value, error)

	value, error = function.MOperations(function.DIV, 1, 2, 3)
	fmt.Println(value, error)

	value, error = function.MOperations(function.MUL, 1, 2, 3)
	fmt.Println(value, error)

	fmt.Println()

	fn := function.FactoryOperation(function.SUM)
	value = fn(1, 2)

	fmt.Println(value)

	fn = function.FactoryOperation(function.SUB)
	value = fn(1, 2)

	fmt.Println(value)
	fn = function.FactoryOperation(function.DIV)
	value = fn(1, 2)

	fmt.Println(value)
	fn = function.FactoryOperation(function.MUL)
	value = fn(1, 2)

	fmt.Println(value)

}

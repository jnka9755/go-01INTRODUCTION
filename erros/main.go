package main

import (
	"errors"
	"fmt"

	"github.com/jnka9755/go-01INTRODUCTION/erros/operator"
)

func main() {

	var err error
	err = errors.New("My new error")
	fmt.Println(err)
	fmt.Println(err.Error())
	fmt.Println()

	err2 := fmt.Errorf("My format error, string: %s, number: %d", "my string", 23)
	fmt.Println(err2)
	fmt.Println(err2.Error())
	fmt.Println()

	defer func() {
		fmt.Println("In main defer")
		r := recover()
		if r != nil {
			fmt.Println("Recovered in ", r)
		}
	}()

	z := operator.Div(4, 0)

	fmt.Println("Results")
	fmt.Println("Z is: ", z)
}

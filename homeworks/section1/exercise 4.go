package main

import (
	"fmt"

	"github.com/jnka9755/go-01INTRODUCTION/homeworks/section1/matrix"
)

func main() {
	m, err := matrix.New([]float64{2, 7, 8}, []float64{4, 4, 7}, []float64{5, 6, 1})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	m.Print()
}

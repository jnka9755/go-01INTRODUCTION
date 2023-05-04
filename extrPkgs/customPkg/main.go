package main

import "github.com/jnka9755/go-01INTRODUCTION/homeworks/section1/matrix"

func main() {

	input, _ := matrix.New(
		[]float64{2, 2, 2},
		[]float64{3, 3, 3},
		[]float64{1, 2, 3},
		[]float64{7, 9, 5},
		[]float64{5, 1, 3},
	)

	input.Print()
}

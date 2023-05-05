package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Number interface {
	int | int32 | int64 | float32 | float64
}

type N1 interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

type (
	MyGenericStruct[D any] struct {
		StrValue string
		Data     D
	}

	MyFirstData struct{}

	MySecondData struct{}
)

type CustomSlice[V int | string] []V
type Point int

func main() {

	v1 := []float64{1.3, 5.43, 12.231, 8.895, 99.12}
	v2 := []int32{9, 23, 1, 12, 6, 50}

	fmt.Println(Sum1(v1))
	fmt.Println(Sum1(v2))
	fmt.Println()

	fmt.Println(Sum2(v1))
	fmt.Println(Sum2(v2))

	fmt.Println()

	anyType1(1, 1)
	anyType1("1", "1")
	fmt.Println()

	anyType2(1, "1")
	fmt.Println()

	comparableType(1, 1)
	fmt.Println()

	orderedValues(2, 4)
	fmt.Println()

	csInt := CustomSlice[int]{1, 2, 3, 4, 5}
	fmt.Println(csInt)

	csStg := CustomSlice[string]{"a", "e", "i", "o", "u"}
	fmt.Println(csStg)

	fmt.Println()

	x, y := Point(5), Point(2)
	v := MinNumber(x, y)
	println(v)
	fmt.Println()

	fd1 := MyGenericStruct[MyFirstData]{StrValue: "Test", Data: MyFirstData{}}
	fd1.Data.PrintOne()

	fd2 := MyGenericStruct[MySecondData]{StrValue: "Test", Data: MySecondData{}}
	fd2.Data.PrintTwo()
}

func Sum1[N int | int32 | int64 | float32 | float64](v []N) N {

	var total N

	for _, item := range v {
		total += item
	}

	return total
}

func Sum2[N Number](v []N) N {

	var total N

	for _, item := range v {
		total += item
	}

	return total
}

func anyType1[N any](v1, v2 N) {
	fmt.Println(v1, v2)
}

func anyType2[N1 any, N2 any](v1 N1, v2 N2) {
	fmt.Println(v1, v2)
}

func comparableType[N comparable](v1, v2 N) {
	fmt.Println(v1, v2)
	fmt.Println(v1 == v2)
}

func orderedValues[N constraints.Ordered](v1, v2 N) {
	fmt.Println(v1, v2)
	fmt.Println(v1 == v2)
	fmt.Println(v1 > v2)
	fmt.Println(v1 < v2)
}

func MinNumber[T N1](x, y T) T {

	if x < y {
		return x
	}

	return y
}

func (d MyFirstData) PrintOne() {
	fmt.Println("Print first")
}

func (d MySecondData) PrintTwo() {
	fmt.Println("Print second")
}

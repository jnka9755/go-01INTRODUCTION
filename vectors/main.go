package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var intVar int
	fmt.Println(intVar)
	fmt.Printf("Type: %T, bytes: %d, bits: %d\n", intVar, unsafe.Sizeof(intVar), unsafe.Sizeof(intVar)*8)

	var intArray [5]int
	fmt.Println(intArray)

	intArray[1] = 2
	intArray[2] = 5
	intArray[3] = 9

	fmt.Println(intArray)

	stringArray := [3]string{"value1", "value2", "value3"}
	fmt.Println(stringArray)
	fmt.Println("Size", len(stringArray))
	fmt.Printf("Type: %T, bytes: %d, bits: %d\n", stringArray, unsafe.Sizeof(stringArray), unsafe.Sizeof(stringArray)*8)

	var intSlice []int
	fmt.Printf("Size: %d, value: %v\n", len(intSlice), intSlice)

	intSlice = append(intSlice, 5, 10, 15, 20, 25, 30)
	fmt.Printf("Size: %d, value: %v\n", len(intSlice), intSlice)

	itemsIntVar := intArray[2:4]
	itemsIntVar1 := intArray[:4]
	itemsIntVar2 := intArray[1:]

	fmt.Println("Items: ", itemsIntVar)
	fmt.Println("Items: ", itemsIntVar1)
	fmt.Println("Items: ", itemsIntVar2)

	slice := make([]int, 3)
	fmt.Println(slice)
	fmt.Printf("Size: %d, value: %v\n", len(slice), slice)
}

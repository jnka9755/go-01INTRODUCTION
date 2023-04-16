package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {

	var intVar int // Positivos y negativos
	intVar = -5
	fmt.Println("My int: ", intVar)

	var uIntVar uint // Solo positivos
	uIntVar = 15
	fmt.Println("My uint: ", uIntVar)

	var stringVar string // Cadenas de texto
	stringVar = "Juan Camilo"
	fmt.Println("My string: ", stringVar)

	var boolVar bool // Valores booleanos
	boolVar = true
	fmt.Println("My bool: ", boolVar)

	fmt.Println("My memory address: ", &intVar) // '&' previo en la variable para obtener la dirección de memoria

	intVar2 := 20
	fmt.Println("My intVar2: ", intVar2) // Instanciar una variable y asignarse a la vez

	stringVar2 := "Instancia rapida y asignación de variable"
	fmt.Println("My string2: ", stringVar2)

	const intConst int = 60 // 'const' para definir una constante
	fmt.Println("My intConst: ", intConst)

	fmt.Println()

	var int8Var int8
	fmt.Printf("Int default value: %d\n", int8Var)
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", int8Var, unsafe.Sizeof(int8Var), unsafe.Sizeof(int8Var)*8)

	var int16Var int16
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", int16Var, unsafe.Sizeof(int16Var), unsafe.Sizeof(int16Var)*8)

	var int32Var int32
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", int32Var, unsafe.Sizeof(int32Var), unsafe.Sizeof(int32Var)*8)

	var int64Var int64
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", int64Var, unsafe.Sizeof(int64Var), unsafe.Sizeof(int64Var)*8)

	var intVar3 int
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", intVar3, unsafe.Sizeof(intVar3), unsafe.Sizeof(intVar3)*8)

	var floatVar1 float32
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", floatVar1, unsafe.Sizeof(floatVar1), unsafe.Sizeof(floatVar1)*8)

	var floatVar2 float64
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", floatVar2, unsafe.Sizeof(floatVar2), unsafe.Sizeof(floatVar2)*8)

	{
		fmt.Println()
		floatVar := 33.11
		fmt.Printf("Type: %T, value: %f\n", floatVar, floatVar)
		floatStrVar := fmt.Sprintf("%.2f", floatVar) //Float to string
		fmt.Printf("Type: %T, value: %s\n", floatStrVar, floatStrVar)

		intVar := 22
		fmt.Printf("Type: %T, value: %d\n", intVar, intVar)
		intStrVar := fmt.Sprintf("%d", intVar) // Int to string
		fmt.Printf("Type: %T, value: %s\n", intStrVar, intStrVar)

		uintVar, err := strconv.ParseInt("55", 0, 8) //String to Int8
		fmt.Println(err)
		fmt.Printf("Type: %T, value: %d\n", uintVar, uintVar)

		uintVar2, err := strconv.ParseInt("5aa5", 0, 8) //String to Int8 (Error)
		fmt.Println(err)
		fmt.Printf("Type: %T, value: %d\n", uintVar2, uintVar2)

		floatVar2, _ := strconv.ParseFloat("-11.55", 64) //String to float64
		fmt.Printf("Type: %T, value: %f\n", floatVar2, floatVar2)
	}
}

package main

import "fmt"

type MyStruct struct {
	ID   int
	Name string
}

func (m MyStruct) Set(name string) {
	fmt.Printf("Addrs: %p\n", &m)
	m.Name = name
}

func (m *MyStruct) SetP(name string) {
	fmt.Printf("Addrs: %p\n", m)
	m.Name = name
}

func main() {

	var i int
	var iP *int

	i = 34
	iP = &i

	fmt.Println(&i)
	fmt.Println(iP)
	fmt.Println(i)
	fmt.Println(*iP)

	fmt.Println()

	*iP = 1
	fmt.Printf("Var i: %d, Pointer i: %d\n", i, *iP)
	fmt.Println()

	myVar := 30
	fmt.Printf("Addrs: %p, Values: %v\n", &myVar, myVar)
	increment(myVar)
	fmt.Println(myVar)
	incrementP(&myVar)
	fmt.Println(myVar)
	fmt.Println()

	var mySlice []int
	mySlice = append(mySlice, 3, 4, 2)
	fmt.Printf("Addrs: %p, Values: %v\n", mySlice, mySlice)
	fmt.Printf("Addrs 1: %p, Addrs 2: %p, Addrs 3: %p\n", &mySlice[0], &mySlice[1], &mySlice[2])

	upateSlice(mySlice)
	fmt.Println(mySlice)
	fmt.Println()

	myStruct := &MyStruct{ID: 123, Name: "Test"}
	fmt.Printf("Addrs: %p\n", myStruct)
	fmt.Printf("Id: %d, Name: %s\n", myStruct.ID, myStruct.Name)
	updateStruct(myStruct)
	fmt.Printf("Id: %d, Name: %s\n", myStruct.ID, myStruct.Name)
	fmt.Println()
	fmt.Printf("Addrs: %p\n", myStruct)
	myStruct.Set("test method")
	fmt.Printf("Id: %d, Name: %s\n", myStruct.ID, myStruct.Name)
	myStruct.SetP("test method 2")
	fmt.Printf("Id: %d, Name: %s\n", myStruct.ID, myStruct.Name)

}

func increment(val int) {
	fmt.Println(&val)
	val++
}

func incrementP(val *int) {
	fmt.Println(val)
	*val++
}

func upateSlice(v []int) {
	fmt.Printf("Addrs: %p\n", v)
	fmt.Printf("Addrs 1: %p, Addrs 2: %p, Addrs 3: %p\n", &v[0], &v[1], &v[2])
	v[0] = 12
}

func updateStruct(v *MyStruct) {

	fmt.Printf("Addrs in function: %p\n", v)
	v.ID = 999
	v.Name = "Update Struct"
}

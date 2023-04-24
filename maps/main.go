package main

import "fmt"

func main() {
	map1 := make(map[int]string)

	map1[1] = "A"
	map1[2] = "B"
	map1[3] = "C"

	fmt.Println(map1)
	fmt.Println(map1[1])

	delete(map1, 2)

	fmt.Println(map1)

	map1[1] = "D"
	fmt.Println(map1)

	map1[7] = ""
	fmt.Println(map1[7])
	fmt.Println(map1[99])

	v1, ok1 := map1[7]
	fmt.Println("The value: ", v1, "Present?", ok1)

	v2, ok2 := map1[99]
	fmt.Println("The value: ", v2, "Present?", ok2)

	map2 := map[int]string{
		1: "Z",
		2: "U",
		3: "W",
	}

	fmt.Println(map2)
}

package main

import "fmt"

func main() {

	sum := 0

	for i := 0; i < 10; i++ {
		sum++
	}

	fmt.Println(sum)

	for sum < 1000 {
		sum++
	}

	fmt.Println(sum)

	for {
		if sum > 1000 {
			break
		}
		sum++
	}

	fmt.Println(sum)

	fmt.Println()
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	for i := range arr {
		fmt.Println("Index:", i, " Value:", arr[i])
	}

	for i, v := range arr {
		fmt.Println("Index:", i, " Value:", v)
	}

	fmt.Println()
	map1 := map[string]float64{
		"A": 12.3,
		"B": 14.11,
		"C": 54.1,
	}

	for k, v := range map1 {
		fmt.Println("Key:", k, " Value: ", v)
	}

	fmt.Println()

	map2 := map[string][]int{
		"X": nil,
		"I": {1, 3, 5, 7, 9},
		"P": {2, 4, 6, 8, 10},
	}

	for k, value := range map2 {
		fmt.Println("Key:", k)
		for _, v := range value {
			fmt.Println("Value:", v)
		}
		fmt.Println()
	}
}

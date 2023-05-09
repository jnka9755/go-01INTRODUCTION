package main

import (
	"fmt"
	"time"
)

func main() {

	go myProcess("A", 15)
	go myProcess("B", 15)

	time.Sleep(3 * time.Second)

	myFirstChannel := make(chan string)

	go myProcessWithChannel("C", myFirstChannel)

	result := <-myFirstChannel
	fmt.Println(result)

	close(myFirstChannel)

	chanelA := make(chan string)
	chanelB := make(chan string)

	go myProcessWithChannel("D", chanelA)
	go myOtherWithChannel("E", chanelB)

	resultA := <-chanelA
	fmt.Println(resultA)

	resultB := <-chanelB
	fmt.Println(resultB)

	close(chanelA)
	close(chanelB)
}

func myProcessWithChannel(p string, c chan string) {
	myProcess(p, 10)
	c <- "OK"
}

func myOtherWithChannel(p string, c chan string) {
	myProcess(p, 20)
	c <- "OK 2"
}

func myProcess(p string, max int) {

	i := 0
	for i < max {
		time.Sleep(150 * time.Millisecond)
		i++
		fmt.Printf("Process: %s - num: %d\n", p, i)
	}
}

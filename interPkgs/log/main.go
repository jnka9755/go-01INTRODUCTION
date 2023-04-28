package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	log.Println("Test")

	// log.Fatal("My error")

	// log.Panic("My error")

	log.Print("Entra")
	date := time.Now()

	file, err := os.Create(fmt.Sprintf("%d.log", date.UnixNano()))
	if err != nil {
		panic(err)
	}

	l := log.New(file, "Success:", log.LstdFlags|log.Lshortfile)
	l.Println("test 1")
	l.Println("test 2")
	l.Println("test 3")
	l.Println("test 4")
	l.Println("test 5")

	l.SetPrefix("Error: ")
	l.Println("test 1")
	l.Println("test 2")
	l.Println("test 3")
	l.Println("test 4")
	l.Println("test 5")
}

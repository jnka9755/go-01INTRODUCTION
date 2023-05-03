package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	fmt.Println(os.Getenv("MY_ENV1"))
	fmt.Println(os.Getenv("MY_ENV2"))
	fmt.Println(os.Getenv("MY_ENV3"))

	if err := godotenv.Load("other/.var"); err != nil {
		fmt.Println(err)
	}

	fmt.Println(os.Getenv("MY_ENV1"))
	fmt.Println(os.Getenv("MY_ENV2"))
	fmt.Println(os.Getenv("MY_ENV3"))
	fmt.Println()

	myEnv, err := godotenv.Read("other/.var")

	fmt.Println(err)
	fmt.Println(myEnv)
	fmt.Println()

	if err := godotenv.Overload(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(os.Getenv("MY_ENV1"))
	fmt.Println(os.Getenv("MY_ENV2"))
	fmt.Println(os.Getenv("MY_ENV3"))
}

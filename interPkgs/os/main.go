package main

import (
	"fmt"
	"os"
)

func main() {

	p := fmt.Println

	file, err := os.Open("myFile.txt")

	if err != nil {
		p("Error :", err)
		os.Exit(1)
	}

	data := make([]byte, 100)
	c, err := file.Read(data)

	if err != nil {
		p("Error :", err)
		os.Exit(1)
	}

	fmt.Printf("Read: %d bytes %q\n", c, data[:c])
	p()

	env1 := os.Getenv("USERNAME")
	p(env1)
	p()

	os.Setenv("MI_ENV", "mi valor")
	p(os.Getenv("MI_ENV"))
	p()

	dir, _ := os.Getwd()
	p(dir)
	p()

	os.Setenv("ENV_EXISTS", "")

	en1 := os.Getenv("ENV_EXISTS")
	en2 := os.Getenv("ENV_NOTEXISTS")

	p(en1)
	p(en2)
	p()

	env, ok := os.LookupEnv("ENV_EXISTS")
	p(env, ok)
	env, ok = os.LookupEnv("ENV_NOTEXISTS")
	p(env, ok)
	p()

	os.Setenv("DB_USERNAME", "jnka")
	os.Setenv("DB_PASSWORD", "12345")
	os.Setenv("DB_HOTS", "127.0.0.1")
	os.Setenv("DB_PORT", "270108")
	os.Setenv("DB_NAME", "users")

	dbURL := os.ExpandEnv("mongodb://${DB_USERNAME}:${DB_PASSWORD}@$DB_HOTS:DB_PORT/DB_NAME")

	p(dbURL)
}

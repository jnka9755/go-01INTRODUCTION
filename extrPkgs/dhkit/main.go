package main

import (
	"encoding/json"
	"fmt"

	"github.com/digitalhouse-tech/go-lib-kit/meta"
	"github.com/digitalhouse-tech/go-lib-kit/response"
)

type User struct {
	FirstName string `myLabel:"lb3"`
	LastName  string `myLabel:"lb4"`
	Email     string `myLabel:"lb2"`
	Age       int    `myLabel:"lb5"`
}

func main() {

	user := &User{
		FirstName: "Juan",
		LastName:  "Santana",
		Email:     "juanka9755@hotmail.com",
		Age:       26,
	}

	print(response.OK("", user, nil, nil))
	print(response.OK("Message", user, nil, nil))

	m := meta.New(1, 5, 15)
	print(response.OK("Message", user, m, nil))

	print(response.Created("Message", user, m, nil))
	print(response.Accepted("Message", user, m, nil))

	print(response.BadRequest("Message error"))
	print(response.NotFound("Message error"))
	print(response.InternalServerError("Message error"))

}

func print(v interface{}) {

	j, _ := json.Marshal(v)
	fmt.Println(string(j))
	fmt.Println()
}

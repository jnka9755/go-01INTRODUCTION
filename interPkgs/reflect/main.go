package main

import (
	"fmt"
	"reflect"
)

type User struct {
	ID        int64    `myLabel:"lb1" myOtherLabel:""lob2`
	Email     string   `myLabel:"lb2"`
	FirstName string   `myLabel:"lb3"`
	LastName  string   `myLabel:"lb4"`
	Age       *float64 `myLabel:"lb5"`
	Address   Address
}

type Address struct {
	Country string
	State   string
}

func main() {

	var age float64 = 26
	user := User{1, "juan@juan.com", "Juan", "Santana", &age, Address{"Colombia", "Mosquera"}}

	Examine(user)
}

func Examine(data interface{}) {

	if reflect.ValueOf(data).Kind() == reflect.Struct {
		v := reflect.ValueOf(data)
		t := reflect.TypeOf(data)

		fmt.Println("Number of fields: ", v.NumField())

		for i := 0; i < v.NumField(); i++ {

			switch v.Field(i).Kind() {
			case reflect.Int, reflect.Int32, reflect.Int64:
				myVar := v.Field(i).Int()
				fmt.Printf("Field:%d Type:%T Value:%v\n", i, myVar, myVar)
			case reflect.String:
				myVar := v.Field(i).String()
				fmt.Printf("Field:%d Type:%T Value:%v\n", i, myVar, myVar)
			case reflect.Ptr:
				fmt.Printf("Field:%d Value:%v\n", i, v.Field(i))
			case reflect.Struct:
				if v.Field(i).Type() == reflect.TypeOf(Address{}) {
					myVar := v.Field(i).Interface().(Address)

					fmt.Println(myVar.Country)
					fmt.Println(myVar.State)
					fmt.Printf("Field:%d Value:%v\n", i, myVar)
				} else {
					fmt.Println("Unsupported type: ", v.Field(i).Type())
				}
			default:
				fmt.Println("Unsupported type: ", v.Field(i).Type())
			}

			c := t.Field(i).Tag

			fmt.Println(c)
			fmt.Println(c.Get("myLabel"))
			fmt.Println(c.Get("myOtherLabel"))
			fmt.Println()
		}
	} else {
		t := reflect.TypeOf(data)
		v := reflect.ValueOf(data)
		k := t.Kind()

		fmt.Println("Type: ", t)
		fmt.Println("Value: ", v)
		fmt.Println("Kind: ", k)
		fmt.Println()
	}
}

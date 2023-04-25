package main

import (
	"encoding/json"
	"fmt"

	"github.com/jnka9755/go-01INTRODUCTION/structs/structsInterface/structs"
	"github.com/jnka9755/go-01INTRODUCTION/structs/structsInterface/vehicles"
)

func main() {

	var p1 structs.Product

	p1.ID = 123
	p1.Name = "PC"

	fmt.Println(p1)
	fmt.Println()

	p2 := structs.Product{}
	p2.ID = 321
	p2.Name = "Celular"

	fmt.Println(p2)
	fmt.Println()

	p3 := structs.Product{
		543,
		"Mouse",
		structs.Type{
			1,
			"abc-123",
			"Tecnologia",
		},
		10,
		150000,
		nil,
	}

	fmt.Println(p3)
	fmt.Println()

	p4 := structs.Product{
		ID:   456,
		Name: "LEGO",
		Type: structs.Type{
			ID:          2,
			Code:        "bcd-411",
			Description: "Juguete",
		},
	}

	fmt.Println(p4)
	fmt.Println()

	p5 := structs.Product{
		ID:   1,
		Name: "PC gamer",
		Type: structs.Type{
			ID:          2,
			Code:        "123-abc",
			Description: "Electrodomestico",
		},
		Tags: []string{"PC", "Comptador", "Ordenador"},
	}

	fmt.Println(p5)
	fmt.Println()

	p6 := structs.Product{
		ID:    1,
		Price: 6500000,
		Name:  "PC gamer",
		Type: structs.Type{
			ID:          2,
			Code:        "123-abc",
			Description: "Electrodomestico",
		},
		Count: 2,
		Tags:  []string{"PC", "Comptador", "Ordenador"},
	}

	v, err := json.Marshal(p6)

	fmt.Println(err)
	fmt.Println(string(v))

	fmt.Println("Precio total: ", p6.TotalPrice())
	p6.SetName("PC")
	p6.AddTags("Tag1", "Tag2", "Tag3")

	fmt.Println(p6)
	fmt.Println()

	p7 := structs.Product{
		ID:    1,
		Price: 100000,
		Name:  "Mouse",
		Type: structs.Type{
			ID:          2,
			Code:        "123-abc",
			Description: "Electrodomestico",
		},
		Count: 2,
		Tags:  []string{"Raton", "Mouse"},
	}

	p8 := structs.Product{
		ID:    1,
		Price: 150000,
		Name:  "Teclado",
		Type: structs.Type{
			ID:          2,
			Code:        "123-abc",
			Description: "Electrodomestico",
		},
		Count: 2,
		Tags:  []string{"Teclado"},
	}

	car := structs.NewCar(1234)
	car.AddProducts(p6, p7, p8)

	fmt.Println("Carrito de compras")
	fmt.Println("Total productos: ", car.TotalProducts())
	fmt.Printf("Total Carrrito: $%.2f\n", car.Total())
	fmt.Println()

	fmt.Println("Vehiculos")
	carV := vehicles.Car{Time: 160}
	fmt.Println(carV.Distance())
	fmt.Println()

	vArray := []string{"Carro", "Moto", "Camion", "Moto", "Camion"}
	var d float64

	for _, v := range vArray {
		fmt.Printf("Vehiculo: %s\n", v)
		veh, err := vehicles.New(v, 400)

		if err != nil {
			fmt.Println("Error: ", err)
			fmt.Println()
			continue
		}

		distance := veh.Distance()
		fmt.Printf("Distance: %.2f\n", distance)
		fmt.Println()
		d += distance
	}
	fmt.Printf("Total distance: %.2f\n", d)
	fmt.Println()
}

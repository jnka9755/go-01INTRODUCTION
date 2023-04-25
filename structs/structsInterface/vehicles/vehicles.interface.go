package vehicles

import (
	"fmt"
)

type Vehicle interface {
	Distance() float64
}

const (
	CarVehicle   = "Carro"
	BikeVehicle  = "Moto"
	TruckVehicle = "Camion"
)

func New(v string, time int) (Vehicle, error) {
	switch v {
	case CarVehicle:
		return &Car{Time: time}, nil
	case BikeVehicle:
		return &Bike{Time: time}, nil
	case TruckVehicle:
		return &Truck{Time: time}, nil
	}
	return nil, fmt.Errorf("Vehiculo '%s' not exists", v)
}

type Car struct {
	Time int
}

func (c *Car) Distance() float64 {
	return 100 * (float64(c.Time) / 60)
}

type Bike struct {
	Time int
}

func (b *Bike) Distance() float64 {
	return 120 * (float64(b.Time) / 60)
}

type Truck struct {
	Time int
}

func (t *Truck) Distance() float64 {
	return 70 * (float64(t.Time) / 60)
}

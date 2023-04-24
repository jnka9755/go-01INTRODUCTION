package function

import (
	"errors"
	"fmt"
)

type Operation int

const (
	SUM Operation = iota
	SUB
	DIV
	MUL
)

func Add(x int, y int) int {
	return x + y
}

func RepeatString(increment int, value string) {
	for i := 0; i < increment; i++ {
		fmt.Println(value)
	}
}

func Calc(op Operation, x, y float64) (float64, error) {
	switch op {
	case SUM:
		return x + y, nil
	case SUB:
		return x - y, nil
	case DIV:
		if y == 0 {
			return 0, errors.New("No es posible dividir por 0")
		}
		return x / y, nil
	case MUL:
		return x * y, nil
	default:
		return 0, errors.New("Ha ocurrido un error")
	}
}

func Split(value float32) (x, y float32) {
	x = value * 4 / 9
	y = value - x

	return
}

func MSum(values ...float32) float32 {

	var sum float32
	for _, v := range values {
		sum += v
	}

	return sum
}

func MOperations(op Operation, values ...float64) (float64, error) {
	if len(values) == 0 {
		return 0, errors.New("No hay valores")
	}

	calc := values[0]

	for _, v := range values[1:] {
		switch op {
		case SUM:
			calc += v
		case SUB:
			calc -= v
		case DIV:
			if v == 0 {
				return 0, errors.New("No es posible dividir por 0")
			}
			calc /= v
		case MUL:
			calc *= v
		default:
			return 0, errors.New("Ha ocurrido un error")
		}
	}
	return calc, nil
}

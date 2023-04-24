package function

func FactoryOperation(op Operation) func(float64, float64) float64 {
	switch op {
	case SUM:
		return sum
	case SUB:
		return sub
	case DIV:
		return div
	case MUL:
		return mul
	}

	return nil
}

func sum(a, b float64) float64 {
	return a + b
}

func sub(a, b float64) float64 {
	return a - b
}

func div(a, b float64) float64 {
	return a / b
}

func mul(a, b float64) float64 {
	return a * b
}

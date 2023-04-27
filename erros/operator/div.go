package operator

func Div(x, y float64) float64 {
	if y == 0 {
		panic("Divisor mustn't be zero")
	}
	return x / y
}

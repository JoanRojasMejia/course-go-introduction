package operator

func Div(x, y float64) float64 {

	if y <= 0 {
		panic("y is less than 0")
	}

	return x / y
}

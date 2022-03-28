package calc

func Compute(x, y float64, fn func(float64, float64) float64) float64 {
	return fn(x, y)
}

func Add(x, y float64) float64 {
	return x + y
}

func Sub(x, y float64) float64 {
	return x - y
}

func Multiply(x, y float64) float64 {
	return x * y
}

func Divide(x, y float64) float64 {
	return x / y
}

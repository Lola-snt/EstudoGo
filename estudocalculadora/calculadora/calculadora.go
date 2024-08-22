package calculadora

type Calculadora struct {
}

func NewCalculadora() Calculadora {
	return Calculadora{}
}

func (c Calculadora) Sum(x, y int) int {
	return x + y
}

func (c Calculadora) Sub(x, y int) int {
	return x - y
}

func (c Calculadora) Div(x, y int) int {
	return x / y
}

func (c Calculadora) Mult(x, y int) int {
	return x * y
}

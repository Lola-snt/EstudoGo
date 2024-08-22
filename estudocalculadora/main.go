package main

import (
	"fmt"
	"github.com/lolasnt/estudogo/estudocalculadora/calculadora"
)

func main() {

	calc := calculadora.NewCalculadora()
	fmt.Printf("soma: %d", calc.Sum(1, 1))
}

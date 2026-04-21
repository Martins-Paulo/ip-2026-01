package main

import (
	"fmt"
	"math"
)

func main() {
	soma := 0.0
	sinal := 1.0

	for i := 1; i <= 101; i += 2 {
		soma += sinal * (1.0 / math.Pow(float64(i), 3))
		sinal *= -1
	}

	pi := math.Cbrt(soma * 32)

	fmt.Println("Valor aproximado de pi:", pi)
}
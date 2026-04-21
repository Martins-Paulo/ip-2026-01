package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Print("Digite x: ")
	fmt.Scan(&x)

	cosSerie := 0.0
	sinal := 1.0
	fat := 1.0
	pot := 1.0

	for i := 0; i < 20; i++ {
		if i > 0 {
			pot *= x * x
			fat *= float64((2*i - 1) * (2 * i))
		}
		cosSerie += sinal * (pot / fat)
		sinal *= -1
	}

	cosReal := math.Cos(x)
	diferenca := math.Abs(cosSerie - cosReal)

	fmt.Println("Cos pela série:", cosSerie)
	fmt.Println("Cos real:", cosReal)
	fmt.Println("Diferença:", diferenca)
}
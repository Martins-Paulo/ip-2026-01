package main

import "fmt"

func main() {
	soma := 0.0
	num := 1.0
	sinal := 1.0

	for i := 15; i >= 1; i-- {
		den := float64(i * i)
		soma += sinal * (num / den)

		num *= 2
		sinal *= -1
	}

	fmt.Println("Resultado:", soma)
}

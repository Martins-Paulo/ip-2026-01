package main

import "fmt"

func main() {
	var x float64

	fmt.Print("Digite X: ")
	fmt.Scan(&x)

	s := x
	fatorial := 1.0
	sinal := -1.0

	for i := 1; i <= 20; i++ {
		fatorial *= float64(i)
		s += sinal * (x / fatorial)
		sinal *= -1
	}

	fmt.Println("Resultado:", s)
}

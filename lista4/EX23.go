package main

import "fmt"

func main() {
	var n int
	fmt.Print("Digite N: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Valor inválido")
		return
	}

	soma := 0.0
	num := 1000.0
	sinal := 1.0

	for i := 1; i <= n; i++ {
		soma += sinal * (num / float64(i))
		num -= 3
		sinal *= -1
	}

	fmt.Println("Resultado:", soma)
}

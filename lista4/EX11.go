package main

import "fmt"

func main() {
	var n int

	fmt.Print("Digite um número: ")
	fmt.Scan(&n)

	if n < 0 {
		fmt.Println("Valor inválido")
		return
	}

	fatorial := 1

	for i := 1; i <= n; i++ {
		fatorial *= i
	}

	fmt.Println("Fatorial:", fatorial)
}

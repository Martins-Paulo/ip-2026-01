package main

import "fmt"

func main() {
	var n int

	fmt.Print("Digite um número: ")
	fmt.Scan(&n)

	ehTriangular := false

	for i := 1; i*(i+1)*(i+2) <= n; i++ {
		if i*(i+1)*(i+2) == n {
			ehTriangular = true
			break
		}
	}

	if ehTriangular {
		fmt.Println("É número triangular")
	} else {
		fmt.Println("Não é número triangular")
	}
}

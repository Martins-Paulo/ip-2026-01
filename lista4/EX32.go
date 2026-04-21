package main

import "fmt"

func main() {
	var n1, n2 int
	fmt.Print("Digite N1: ")
	fmt.Scan(&n1)
	fmt.Print("Digite N2: ")
	fmt.Scan(&n2)

	resultado := 0

	for i := 0; i < n2; i++ {
		resultado += n1
	}

	fmt.Println("Resultado:", resultado)
}
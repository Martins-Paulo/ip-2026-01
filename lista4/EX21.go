package main

import "fmt"

func main() {
	var b, n int
	resultado := 1

	fmt.Print("Digite b: ")
	fmt.Scan(&b)

	fmt.Print("Digite n: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		resultado *= b
	}

	fmt.Println("Resultado:", resultado)
}

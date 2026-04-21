package main

import "fmt"

func main() {
	var n int

	fmt.Print("Digite N (>=3): ")
	fmt.Scan(&n)

	a, b := 0, 1

	fmt.Print(a, " ", b, " ")

	for i := 3; i <= n; i++ {
		c := a + b
		fmt.Print(c, " ")

		a = b
		b = c
	}
}

package main

import "fmt"

func main() {
	var soma float64 = 0

	num1 := 37.0
	num2 := 38.0

	for i := 1; i <= 37; i++ {
		soma += (num1 * num2) / float64(i)
		num1--
		num2--
	}

	fmt.Println("Soma:", soma)
}

package main

import "fmt"

func main() {
	h := 0.0
	num := 1.0

	for i := 1; i <= 50; i++ {
		h += num / float64(i)
		num += 2
	}

	fmt.Println("Valor de H:", h)
}

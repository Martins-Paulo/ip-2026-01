package main

import "fmt"

func main() {
	var total uint64 = 0
	var graos uint64 = 1

	for i := 1; i <= 64; i++ {
		total += graos
		graos *= 2
	}

	fmt.Println("Total de grãos:", total)
}

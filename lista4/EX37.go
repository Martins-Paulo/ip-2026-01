package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	base := 1
	decimal := 0

	for n > 0 {
		dig := n % 10
		decimal += dig * base
		base *= 8
		n /= 10
	}

	fmt.Println("Decimal:", decimal)
}
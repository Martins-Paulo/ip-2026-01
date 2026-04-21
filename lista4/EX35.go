package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	bin := ""

	for n > 0 {
		bin = fmt.Sprintf("%d", n%2) + bin
		n /= 2
	}

	fmt.Println("Binário:", bin)
}
package main

import "fmt"

func main() {
	var n int
	fmt.Print("Digite N: ")
	fmt.Scan(&n)

	soma := 0
	for i := 1; i <= n; i++ {
		soma += i
	}

	fmt.Println("Soma:", soma)
}
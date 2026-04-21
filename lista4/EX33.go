package main

import "fmt"

func main() {
	var n1, n2 int
	fmt.Scan(&n1, &n2)

	quociente := 0
	resto := n1

	for resto >= n2 {
		resto -= n2
		quociente++
	}

	fmt.Println("Quociente:", quociente)
	fmt.Println("Resto:", resto)
}

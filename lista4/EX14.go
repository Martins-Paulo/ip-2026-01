package main

import "fmt"

func main() {
	var n1, n2 int

	fmt.Print("Digite N1: ")
	fmt.Scan(&n1)

	fmt.Print("Digite N2: ")
	fmt.Scan(&n2)

	for i := n1; i <= n2; i++ {
		if i < 2 {
			continue
		}

		primo := true

		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				primo = false
				break
			}
		}

		if primo {
			fmt.Print(i, " ")
		}
	}
}

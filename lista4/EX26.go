package main

import "fmt"

func main() {
	soma := 0.0
	fat := 1.0

	for i := 0; i < 20; i++ {
		if i > 0 {
			fat *= float64(i)
		}
		soma += float64(100-i) / fat
	}

	fmt.Println("Resultado:", soma)
}
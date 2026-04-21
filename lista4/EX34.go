package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	mmc := a
	for mmc%b != 0 {
		mmc += a
	}

	fmt.Println("MMC:", mmc)
}

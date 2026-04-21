package main

import (
	"fmt"
	"math"
)

func main() {
	for r := 0.0; r <= 20.0; r += 0.5 {
		volume := (4.0 / 3.0) * math.Pi * math.Pow(r, 3)
		fmt.Printf("R=%.1f -> Volume=%.2f\n", r, volume)
	}
}
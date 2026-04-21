package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	hex := ""
	digitos := "0123456789ABCDEF"

	for n > 0 {
		hex = string(digitos[n%16]) + hex
		n /= 16
	}

	fmt.Println("Hexadecimal:", hex)
}
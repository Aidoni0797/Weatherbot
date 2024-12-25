package main

import (
	"fmt"
)

func main() {
	// Ввод катетов
	var a, b, aa, bb, cc int
	fmt.Scan(&a)
	b = a % 1000
	aa = (b % 1000) / 100
	bb = (b % 100) / 10
	cc = b % 10
	fmt.Println(aa + bb + cc)
}

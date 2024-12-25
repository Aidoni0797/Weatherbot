package main

import (
	"fmt"
)

func main() {
	// Ввод катетов
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	fmt.Println(a*b*c, a+b+c)
}

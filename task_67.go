package main

import (
	"fmt"
)

func main() {
	// Ввод катетов
	var a int
	fmt.Scan(&a)

	fmt.Println((a % 1000) / 100)
}

package main

import "fmt"

func main() {
	// Ввод целого числа
	var num int
	fmt.Scanln(&num)

	fmt.Print(num % 10)
}

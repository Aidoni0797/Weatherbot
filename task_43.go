package main

import "fmt"

func main() {
	// Ввод целого числа
	var num int
	fmt.Scanln(&num)

	// Вывод предыдущего и следующего числа с пояснением
	fmt.Print(num % 10)
}

package main

import "fmt"

func main() {
	// Ввод целого числа
	var num int
	fmt.Scanln(&num)

	// Вывод предыдущего и следующего числа с пояснением
	fmt.Printf("Следующее за числом %d число: %d\n", num, num+1)
	fmt.Printf("Для числа %d предыдущее число: %d", num, num-1)
}

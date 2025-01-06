package main

import "fmt"

func main() {
	var numbers [6]int
	numbers[0] = 5 // записать в первый элемент массива 5
	numbers[1] = 2
	numbers[4] = 3
	numbers[5] = 2          // ошибка, индекс вне границ массива
	fmt.Println(numbers[0]) // вывести первый элемент массива
}

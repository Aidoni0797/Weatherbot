package main

import "fmt"

func main() {
	var size, cnt, cnt1 int
	fmt.Scan(&size)                     // считываем сколько чисел будем вводить с клавиатуры
	numbers := make([]int, size)        // заводим новый slice целого типа размерностью кол-ва чисел, которые будем вводить с клавиатуры
	for i := 0; i < len(numbers); i++ { // получаем длину slice-а и проходим по нему циклом
		fmt.Scan(&numbers[i])
	}
	for i := 0; i < len(numbers); i++ { // получаем длину slice-а и проходим по нему циклом
		cnt1 = numbers[i]
		if (i > 0) && (i%2 != 0) {
			cnt = numbers[i-1]
			numbers[i-1] = cnt1
			numbers[i] = cnt
		}
	}
	for i := 0; i < len(numbers); i++ {
		fmt.Print(numbers[i], " ")
	}
}

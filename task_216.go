package main

import "fmt"

func main() {
	var size, cnt int
	fmt.Scan(&size)                     // считываем сколько чисел будем вводить с клавиатуры
	numbers := make([]int, size)        // заводим новый slice целого типа размерностью кол-ва чисел, которые будем вводить с клавиатуры
	for i := 0; i < len(numbers); i++ { // получаем длину slice-а и проходим по нему циклом
		fmt.Scan(&numbers[i])
		if i > 0 {
			if numbers[i] > numbers[i-1] {
				cnt = cnt + 1
			}
		}
	}
	fmt.Print(cnt)
}

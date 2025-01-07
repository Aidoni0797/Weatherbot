package main

import "fmt"

func main() {
	var size, cnt int
	fmt.Scan(&size)                     // считываем сколько чисел будем вводить с клавиатуры
	numbers := make([]int, size)        // заводим новый slice целого типа размерностью кол-ва чисел, которые будем вводить с клавиатуры
	for i := 0; i < len(numbers); i++ { // получаем длину slice-а и проходим по нему циклом
		fmt.Scan(&numbers[i])
	}
	for i := 0; i < len(numbers); i++ { // получаем длину slice-а и проходим по нему циклом
		if i > 0 {
			if (numbers[i] > 0 && numbers[i-1] > 0) || (numbers[i] < 0 && numbers[i-1] < 0) {
				fmt.Print("YES")
				cnt = -1
				break
			}
		}

	}
	if cnt != -1 {
		fmt.Print("NO")
	}
}

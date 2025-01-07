package main

import "fmt"

func main() {
	var size int
	fmt.Scan(&size)                     // считываем сколько чисел будем вводить с клавиатуры
	numbers := make([]int, size)        // заводим новый slice целого типа размерностью кол-ва чисел, которые будем вводить с клавиатуры
	for i := 0; i < len(numbers); i++ { // получаем длину slice-а и проходим по нему циклом
		fmt.Scan(&numbers[i])
		if i%3 == 0 {
			fmt.Print(numbers[i], " ")
		} // считываем и сразу записываем значение в i-ый индекс slice-а
	}

}

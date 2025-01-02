package main

import "fmt"

func main() {
	var num, result int
	for {
		fmt.Scan(&num) // Считываем число

		// Пропускаем числа меньше 10
		if num < 10 {
			continue
		}
		if num >= 10 && num <= 100 {
			// Выводим число, если оно от 10 до 100 (включительно)
			fmt.Println(result)
		}
		// Прерываем цикл, если число больше 100
		if num > 100 {
			break
		}

	}

}

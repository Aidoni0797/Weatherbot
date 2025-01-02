package main

import "fmt"

func main() {
	var num, sum int

	for {
		fmt.Scan(&num) // Считываем число
		if num == 0 {
			break // Завершаем, если встречаем 0
		}

		// Проверяем условия: число кратно 2, но не кратно 3
		if num%2 == 0 && num%3 != 0 {
			sum += num
		}
	}

	fmt.Println(sum) // Выводим результат
}

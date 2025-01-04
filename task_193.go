package main

import "fmt"

func main() {
	var num int
	max1, max2 := 0, 0

	// Считываем последовательность
	for {
		fmt.Scan(&num)
		if num == 0 {
			break
		}

		// Обновляем максимумы
		if num > max1 {
			max2 = max1
			max1 = num
		} else if num > max2 {
			max2 = num
		}
	}

	// Вывод результата
	fmt.Println(max2)
}

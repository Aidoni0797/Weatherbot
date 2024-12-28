package main

import "fmt"

func main() {
	var n int
	// Ввод числа
	fmt.Scan(&n)

	// Извлекаем цифры числа
	digit1 := n / 100       // Первая цифра
	digit2 := (n / 10) % 10 // Вторая цифра
	digit3 := n % 10        // Третья цифра

	// Проверка на возрастающую последовательность
	if digit1 < digit2 && digit2 < digit3 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

package main

import "fmt"

func main() {
	var n int
	// Ввод числа
	fmt.Scan(&n)

	// Извлекаем цифры числа
	digit1 := n / 1000       // Первая цифра
	digit2 := (n / 100) % 10 // Вторая цифра
	digit3 := (n / 10) % 10  // Третья цифра
	digit4 := n % 10         // Четвертая цифра

	// Проверяем симметричность
	if digit1 == digit4 && digit2 == digit3 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

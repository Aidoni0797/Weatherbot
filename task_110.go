package main

import "fmt"

func main() {
	var number int
	// Ввод числа
	fmt.Scan(&number)

	// Извлекаем последнюю цифру числа
	lastDigit := number % 10

	// Проверяем, чётная ли последняя цифра
	if lastDigit%2 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

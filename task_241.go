package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m) // Считываем размеры массива

	// Создаём таблицу умножения
	for i := 1; i <= n; i++ { // Нумерация строк с 1
		for j := 1; j <= m; j++ { // Нумерация столбцов с 1
			fmt.Print(i*j, " ") // Выводим элемент таблицы
		}
		fmt.Println() // Переход на новую строку
	}
}

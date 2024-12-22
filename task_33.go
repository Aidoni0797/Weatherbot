package main

import "fmt"

func main() {
	// Ввод данных
	var a, b, n int
	fmt.Scan(&a, &b, &n)

	// Рассчитываем общую стоимость в копейках
	totalCents := (a*100 + b) * n

	// Переводим в рубли и копейки
	rubles := totalCents / 100
	cents := totalCents % 100

	// Вывод результата
	fmt.Println(rubles, cents)
}

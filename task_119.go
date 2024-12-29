package main

import (
	"fmt"
)

func main() {
	var class1, class2, class3 int

	// Ввод количества учащихся в каждом из трех классов
	fmt.Scan(&class1, &class2, &class3)

	// Вычисляем количество парт для каждого класса
	desks1 := (class1 + 1) / 2
	desks2 := (class2 + 1) / 2
	desks3 := (class3 + 1) / 2

	// Суммируем общее количество парт
	totalDesks := desks1 + desks2 + desks3

	// Выводим результат
	fmt.Println(totalDesks)
}

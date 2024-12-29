package main

import (
	"fmt"
)

func main() {
	var techCount int

	// Считываем количество пройденных технологий
	fmt.Scan(&techCount)

	// Определяем уровень программиста
	switch {
	case techCount <= 3:
		fmt.Println("начинающий")
	case techCount >= 4 && techCount <= 7:
		fmt.Println("младший разработчик")
	case techCount >= 8 && techCount <= 15:
		fmt.Println("средний разработчик")
	case techCount >= 16:
		fmt.Println("старший разработчик")
	default:
		fmt.Println("Некорректные данные")
	}
}

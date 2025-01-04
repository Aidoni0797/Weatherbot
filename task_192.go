package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	// Инициализация переменных
	maxSpeed := 0
	belowThirty := false

	// Считывание скоростей
	for i := 0; i < n; i++ {
		var speed int
		fmt.Scan(&speed)

		// Обновление максимальной скорости
		if speed > maxSpeed {
			maxSpeed = speed
		}

		// Проверка на скорость меньше 30
		if speed < 30 {
			belowThirty = true
		}
	}

	// Вывод результата
	if belowThirty {
		fmt.Printf("%d YES\n", maxSpeed)
	} else {
		fmt.Printf("%d NO\n", maxSpeed)
	}
}

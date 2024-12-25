package main

import (
	"fmt"
)

func main() {
	// Ввод количества бит
	var bits int
	fmt.Scan(&bits)

	// Перевод бит в килобайты
	kilobytes := float64(bits) / 8 / 1024

	// Вывод результата
	fmt.Println(kilobytes)
}

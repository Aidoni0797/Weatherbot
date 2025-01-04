package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n string
	fmt.Scan(&n)

	// Удаляем цифры '5' и '7' из числа
	result := strings.ReplaceAll(n, "5", "")
	result = strings.ReplaceAll(result, "7", "")

	// Преобразуем строку в число, чтобы убрать ведущие нули
	if result == "" {
		result = "0" // Если все цифры были удалены, выводим 0
	} else {
		// Преобразуем в число и обратно в строку, чтобы убрать ведущие нули
		num, _ := strconv.Atoi(result)
		result = strconv.Itoa(num)
	}

	// Выводим результат
	fmt.Println(result)
}

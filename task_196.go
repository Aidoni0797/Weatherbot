package main

import (
	"fmt"
	"strings"
)

func main() {
	var n string
	fmt.Scan(&n)

	// можно оказывается и так а я думаю всех делить на 10 потом сравнить с 5 и 7 в том числе потом еще что то, хех Chatgpt ты умный все такий
	// Удаляем цифры '5' и '7' из числа
	result := strings.ReplaceAll(n, "5", "")
	result = strings.ReplaceAll(result, "7", "")

	// Вывод результата
	fmt.Println(result)
}

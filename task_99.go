package main

import (
	"fmt"
)

func main() {
	var weight int

	// Ввод веса боксера
	fmt.Scan(&weight)

	// Определение категории
	if weight < 60 {
		fmt.Println("Легкий вес")
	} else if weight < 64 {
		fmt.Println("Первый полусредний вес")
	} else if weight < 69 {
		fmt.Println("Полусредний вес")
	}
}

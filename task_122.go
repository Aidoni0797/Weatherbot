package main

import "fmt"

func main() {
	var x int

	// Ввод числа
	fmt.Scan(&x)

	// Проверяем принадлежность числу промежуткам: [-30, -2] или [7, 25]
	if (x >= -30 && x <= -2) || (x >= 7 && x <= 25) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

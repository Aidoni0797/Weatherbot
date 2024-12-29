package main

import (
	"fmt"
)

func main() {
	var a, b, c int

	// Ввод данных
	fmt.Scan(&a, &b, &c)

	// Проверка существования невырожденного треугольника
	if a+b > c && a+c > b && b+c > a {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

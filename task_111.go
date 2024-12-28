package main

import "fmt"

func main() {
	var n int
	// Ввод числа
	fmt.Scan(&n)

	// Если число чётное, прибавляем 2, иначе прибавляем 1
	if n%2 == 0 {
		fmt.Println(n + 2)
	} else {
		fmt.Println(n + 1)
	}
}

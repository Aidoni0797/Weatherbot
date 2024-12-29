package main

import "fmt"

func main() {
	var a, b int

	// Ввод данных
	fmt.Scan(&a, &b)

	for i := a; i <= b; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

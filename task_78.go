package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	if a >= 12 {
		fmt.Println("Доступ разрешен")
	} else {
		fmt.Println("Доступ запрещен")
	}
}

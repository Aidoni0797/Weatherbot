package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n > 0 {
		fmt.Print("Положительное")
	} else {
		if n < 0 {
			fmt.Print("Отрицательное")
		} else {
			fmt.Print("Ноль")
		}
	}
}

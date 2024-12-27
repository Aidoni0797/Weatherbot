package main

import "fmt"

func main() {
	var a, b string
	fmt.Scan(&a, &b)

	if a == b {
		fmt.Println("Пароль принят")
	} else {
		fmt.Println("Пароль не принят")
	}
}

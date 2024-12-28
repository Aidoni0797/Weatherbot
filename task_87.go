package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	if a <= -3 || a >= 7 {
		fmt.Print("Принадлежит")
	} else {
		fmt.Print("Не принадлежит")
	}
}

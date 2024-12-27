package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	if a > -1 && a < 17 {
		fmt.Print("Принадлежит")
	} else {
		fmt.Print("Не принадлежит")
	}
}

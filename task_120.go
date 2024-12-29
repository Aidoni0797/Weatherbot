package main

import "fmt"

func main() {
	var w int
	fmt.Scan(&w)

	// Проверяем, можно ли разделить арбуз на две четные части
	if w > 2 && w%2 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

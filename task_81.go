package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	if a > 0 {
		fmt.Println(1)
	}
	if a == 0 {
		fmt.Println(1)
	}
	if a < 0 {
		fmt.Println(-1)
	}
}

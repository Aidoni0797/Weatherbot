package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fact := 1

	for i := a; i <= b; i++ {
		fact = fact * i
	}

	fmt.Println(fact)
}

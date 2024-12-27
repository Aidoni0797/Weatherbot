package main

import "fmt"

func main() {
	a := 2
	b := 3
	c := 2
	fmt.Println(a > b)  // false
	fmt.Println(a < b)  // true
	fmt.Println(a >= b) // false
	fmt.Println(a >= c) // true
	fmt.Println(a == b) // false
	fmt.Println(a == c) // true
	fmt.Println(a != b) // true
}

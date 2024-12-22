package main

import "fmt"

func main() {
	x := 432
	y := x / 100
	x = (x % 100) * 10
	x = x + y
	fmt.Println(x)
}

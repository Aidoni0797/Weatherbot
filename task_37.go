package main

import "fmt"

func main() {
	var number int = 1324
	fmt.Println(number % 10)    // 4
	fmt.Println(number % 100)   // 24
	fmt.Println(number % 1000)  // 324
	fmt.Println(number % 10000) // 1324
}

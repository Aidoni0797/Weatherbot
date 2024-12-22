package main

import "fmt"

func main() {
	var number int = 1324
	fmt.Println(number / 10)    // 132
	fmt.Println(number / 100)   // 13
	fmt.Println(number / 1000)  // 1
	fmt.Println(number / 10000) // 0
}

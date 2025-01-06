package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(numbers); i++ {
		fmt.Print(numbers[i], " ") // или fmt.Println(numbers[i])
	}
}

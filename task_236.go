package main

import "fmt"

func main() {
	var count int
	fmt.Scan(&count)
	numbers := make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Scan(&numbers[i])
	}

	cnt := 0
	numbersSize := len(numbers)
	for i, number := range numbers {
		for j := i + 1; j < numbersSize; j++ {
			if number == numbers[j] {
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}

package main

import "fmt"

func main() {
	var size, cnt, cnt1 int
	fmt.Scan(&size)
	numbers := make([]int, size)
	numbers1 := make([]int, size)
	for i := 0; i < len(numbers); i++ {
		fmt.Scan(&numbers[i])
		if i < len(numbers)-1 {
			numbers1[i] = numbers[i]
		}
	}
	cnt = numbers[len(numbers)-1]
	for i := 0; i < len(numbers); i++ {
		if i > 0 {
			cnt1 = numbers1[i-1]
			numbers[i] = cnt1
		}

	}
	numbers[0] = cnt
	fmt.Print(numbers)
}

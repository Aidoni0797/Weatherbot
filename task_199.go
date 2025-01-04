package main

import (
	"fmt"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	arr := []int{}
	var min, cnt int
	var number int

	_, _ = fmt.Scan(&min)

	for i := 1; i < n; i++ {
		_, _ = fmt.Scan(&number)
		arr = append(arr, number)
		if number < min {
			min = number
		}
	}

	for i := 0; i < n-1; i++ {
		if arr[i] == min {
			cnt = cnt + 1
		}
	}

	fmt.Println(cnt)
}

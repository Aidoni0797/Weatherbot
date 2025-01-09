package main

import "fmt"

func main() {
	var count, cnt int
	fmt.Scan(&count)
	numbers := make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Scan(&numbers[i])
	}

	for i := 0; i < count; i++ {
		for j := i + 1; j < count; j++ {
			if numbers[i] == numbers[j] {
				fmt.Print("YES")
				cnt = -1
				break
			}
		}
	}
	if cnt != -1 {
		fmt.Print("NO")
	}
}

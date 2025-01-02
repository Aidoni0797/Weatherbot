package main

import "fmt"

func main() {
	var num, sum int
	fmt.Scan(&num)
	for i := 1; i <= num; i++ {
		if i%10 == 7 {
			sum = sum + 1
		}
	}

	fmt.Println(sum)
}

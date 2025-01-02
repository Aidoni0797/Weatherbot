package main

import "fmt"

func countDivisors(n int) int {
	count := 0
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			count++ // i - делитель
			if i != n/i {
				count++ // n / i - делитель
			}
		}
	}
	return count
}

func main() {
	var a, b, k int
	fmt.Scan(&a, &b, &k)

	for i := a; i <= b; i++ {
		if countDivisors(i) >= k {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}

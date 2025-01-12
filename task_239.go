package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if j == n-i-1 {
				matrix[i][j] = 1
			} else if j > n-i-1 {
				matrix[i][j] = 2
			}
		}
	}

	for _, row := range matrix {
		for _, val := range row {
			fmt.Print(val, " ")
		}
		fmt.Println()
	}
}

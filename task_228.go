package main

import "fmt"

func main() {
	a := []int{1, 2, 18, 8, 14, 9, 23, 7, 51, 99}
	i := 0
	j := 9
	for a[i] < 10 {
		i = i + 1
	}
	for a[j] > 10 {
		j = j - 1
	}
	t := a[i] - a[j]
	fmt.Print(t)
}

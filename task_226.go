package main

import "fmt"

func main() {
	a := []int{3, 14, 15, 92, 6, 2, 7, 18, 28, 17}
	s := 0
	n := 10
	for i := 2; i < n; i++ {
		s = s + a[i] - a[i-2]
	}
	fmt.Print(s)
}

package main

import "fmt"

func main() {
	a := []int{7, 9, 8, 1, 2, 3, 3, 10, 8, 6}
	s := 0
	for i := 1; i < 10; i++ {
		if a[i-1] < a[i] {
			a[i] = a[i-1] + 1
			s = s + 1
		}
	}
	fmt.Print(s)
}

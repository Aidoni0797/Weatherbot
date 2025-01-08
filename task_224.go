package main

import "fmt"

func main() {
	a := []int{5, 1, 6, 7, 8, 8, 7, 7, 6, 9}
	c := 0
	for i := 1; i < 10; i++ {
		if a[i-1] >= a[i] {
			t := a[i]
			a[i] = a[i-1]
			a[i-1] = t
		} else {
			c = c + 1
		}
	}
	fmt.Print(c)
}

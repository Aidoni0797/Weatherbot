package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a)

	b = a % 10
	c = a / 10 % 10
	d = a / 100 % 10

	if b == c || c == d || b == d {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}

}

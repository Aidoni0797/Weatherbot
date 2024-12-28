package main

import "fmt"

func main() {
	var a, b, c, d, e, f, g int
	fmt.Scan(&a)

	b = a % 10
	c = a / 10 % 10
	d = a / 100 % 10
	e = a / 1000 % 10
	f = a / 10000 % 10
	g = a / 100000 % 10

	if b+c+d != e+f+g {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}

}

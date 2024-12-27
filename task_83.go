package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	if x > 5 || y < 10 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}

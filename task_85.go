package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	if a >= 100 && a <= 999 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}

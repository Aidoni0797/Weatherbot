package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	if a%2 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

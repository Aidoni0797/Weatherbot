package main

import "fmt"

func main() {
	var grade int
	fmt.Scan(&grade)
	if grade >= 90 {
		fmt.Print(5)
	}

	if grade >= 80 {
		fmt.Print(4)
	}

	if grade >= 70 {
		fmt.Print(3)
	}

	if grade >= 60 {
		fmt.Print(2)
	}

	if grade >= 0 {
		fmt.Print(1)
	}
}

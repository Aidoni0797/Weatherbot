package main

import (
	"fmt"
)

func main() {
	var a, b, c, d, mest1, mest2 int

	fmt.Scan(&a, &b, &c, &d)

	mest1 = a - c
	mest2 = b - d

	if mest1 < 0 {
		mest1 = -1 * mest1
	}

	if mest2 < 0 {
		mest2 = -1 * mest2
	}

	if (a == c) || (b == d) || (mest1 == mest2) {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}

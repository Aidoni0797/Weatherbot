package main

import (
	"fmt"
)

func main() {
	var a, b, c, d, mest1, mest2, mest3, mest4, mest5, mest6, mest7, mest8 int
	var mest9, mest10, mest11, mest12, mest13, mest14, mest15, mest16 int
	fmt.Scan(&a, &b, &c, &d)

	mest1 = a + 2
	mest2 = b + 1

	mest3 = a + 2
	mest4 = b - 1

	mest5 = a + 1
	mest6 = b + 2

	mest7 = a - 1
	mest8 = b + 2

	mest9 = a - 2
	mest10 = b + 1

	mest11 = a - 2
	mest12 = b - 1

	mest13 = a + 1
	mest14 = b - 2

	mest15 = a - 1
	mest16 = b - 2

	if (c == mest1 && d == mest2) || (c == mest3 && d == mest4) || (c == mest5 && d == mest6) || (c == mest7 && d == mest8) || (c == mest9 && d == mest10) || (c == mest11 && d == mest12) || (c == mest13 && d == mest14) || (c == mest15 && d == mest16) {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}

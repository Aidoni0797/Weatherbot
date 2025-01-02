// код из курса, хех, когда ты сама напишешь такие коды
package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	for x := 1; x <= 10000; x++ {
		if a*x*x+b*x+c == 0 {
			fmt.Println(x)
		}
	}
}

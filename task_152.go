// код из курса ничего себе
package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	pokazatel := 0
	stepen := 1

	for pokazatel < n {
		stepen = stepen * 2
		pokazatel++
	}

	fmt.Println(stepen)
}

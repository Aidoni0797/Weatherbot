// как в жизне, если твои усилия равняется нулю ты же остановишся и все дальше никаких действии нет, больно
package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	for n != 0 { // пока число, хранимое в переменной n не равно нулю
		fmt.Scan(&n)
	}
}
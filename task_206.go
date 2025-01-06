package main

import (
	"fmt"
	"strconv"
)

func main() {
	var iDONi, c int
	var stp string
	fmt.Scan(&iDONi)
	c = 0
	iDONi = iDONi + 1
	for c == 0 {
		stp = strconv.Itoa(iDONi)
		if (stp[0] != stp[1]) && (stp[0] != stp[2]) && (stp[0] != stp[3]) && (stp[1] != stp[2]) && (stp[1] != stp[3]) && (stp[2] != stp[3]) {
			c = -1
			fmt.Print(iDONi)
			break
		} else {
			iDONi = iDONi + 1
		}
	}
}

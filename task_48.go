package main

import "fmt"

func main() {
	a := 5.0 // наличие точки дает понять компилятору что это вещественное число
	b := 2.0

	sum := a + b
	diff := a - b
	mult := a * b

	fmt.Println(a, " + ", b, " = ", sum)
	fmt.Println(a, " - ", b, " = ", diff)
	fmt.Println(a, " * ", b, " = ", mult)
}

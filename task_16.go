package main

import "fmt"

func main() {
	var a float64 = 6.6
	b := int(a)    // приведение (casting) вещественного числа к целому
	fmt.Println(b) // Вывод 6
}

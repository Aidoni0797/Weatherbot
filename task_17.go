package main

import "fmt"

func main() {
	var a int = 6
	b := float64(a)      // приведение (casting) целого числа к вещественному
	fmt.Println(b + 0.4) // Вывод 6.4
}

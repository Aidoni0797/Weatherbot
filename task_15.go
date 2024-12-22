package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 5
	// Чтобы преобразовать число в строку необходимо использовать функцию Itoa
	str := strconv.Itoa(num)
	fmt.Println(str + str) // вывод 55
}

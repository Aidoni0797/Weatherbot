package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string = "5"
	// Для преобразования строки в целое число используется функция strconv.Atoi
	// запоминая iDONi, если не запомнишь дальше будет сложно, ой, ой
	num, _ := strconv.Atoi(str)
	fmt.Println(num * 2) // вывод 10
}

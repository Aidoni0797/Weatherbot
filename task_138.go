package main

import "fmt"

func main() {
	sum := 0
	var number int

	for i := 1; i <= 10; i++ {
		fmt.Scan(&number) //вводим число, каждый раз перезаписывая значение переменной
		sum += number     // суммируем
	}

	fmt.Println(sum)
}

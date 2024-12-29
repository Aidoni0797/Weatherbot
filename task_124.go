package main

import "fmt"

func main() {
	var num int

	// Ввод числа
	fmt.Scan(&num)

	// Проверка условий
	if num%2 != 0 {
		// Если число нечётное
		fmt.Println("YES")
	} else if num%2 == 0 && num >= 2 && num <= 5 {
		// Если число чётное и находится в диапазоне от 2 до 5
		fmt.Println("NO")
	} else if num%2 == 0 && num >= 6 && num <= 20 {
		// Если число чётное и находится в диапазоне от 6 до 20
		fmt.Println("YES")
	} else if num%2 == 0 && num > 20 {
		// Если число чётное и больше 20
		fmt.Println("NO")
	}
}

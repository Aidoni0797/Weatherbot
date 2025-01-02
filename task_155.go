// уважаемый chatgpt, перед тобою приклоняюсь
package main

import "fmt"

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b // Алгоритм Евклида: берем остаток от деления
	}
	return a
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)       // Ввод двух чисел
	fmt.Println(gcd(a, b)) // Вывод НОД
}

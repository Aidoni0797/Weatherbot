package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	// Преобразуем число в строку
	str := strconv.Itoa(n)

	// Проверяем, является ли строка палиндромом
	isPalindrome := true
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			isPalindrome = false
			break
		}
	}

	// Вывод результата
	if isPalindrome {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

package main

import (
	"fmt"
)

func main() {
	// Ввод количества элементов массива
	var N int
	fmt.Scan(&N)

	// Ввод элементов массива
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&arr[i])
	}

	// Проверка, является ли массив палиндромом
	isPalindrome := true
	for i := 0; i < N/2; i++ {
		if arr[i] != arr[N-i-1] {
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

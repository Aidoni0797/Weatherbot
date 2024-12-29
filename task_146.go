package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fact := 1 // переменная, в которой будет хранится итоговое произведение

	for i := 1; i <= n; i++ {
		fact = fact * i // умножаем текущее произведение на число и снова записываем в нее.
	}

	fmt.Println(fact)
}

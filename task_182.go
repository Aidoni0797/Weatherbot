package main

import "fmt"

func main() {
	var n, c, d int
	fmt.Scan(&n)
	fmt.Scan(&c)
	fmt.Scan(&d)

	// Проходим по всем числам от 1 до n
	for i := 1; i <= n; i++ {
		// Если число кратно c, но не кратно d, выводим его
		if i%c == 0 && i%d != 0 {
			fmt.Println(i)
		}
	}
}

package main

import "fmt"

func main() {
	var isTrue bool
	fmt.Scan(&isTrue)

	if isTrue {
		a := 5             // Доступна только внутри данного блока
		fmt.Println(a + 2) // ОК
	} else {
		b := 10            // Доступна только внутри данного блока
		fmt.Println(b + 1) // ОК

		fmt.Println(a) // Ошибка. Такой переменной не существует здесь
	}

	fmt.Println(a) // Ошибка. Такой переменной не существует здесь
	fmt.Println(b) // Ошибка. Такой переменной не существует здесь
}

package main

import "fmt"

func main() {
	fmt.Println("Выйти из дома")
	fmt.Println("Дойти до магазина")
	fmt.Println("Хлеб есть? Введите Есть или Нет")

	var answer string
	fmt.Scan(&answer)

	if answer == "Есть" {
		fmt.Println("Купить хлеб")
	}

	fmt.Println("Дойти до дома")
	fmt.Println("Зайти в дом")
}

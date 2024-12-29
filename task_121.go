package main

import "fmt"

func main() {
	var age int
	var gender string

	// Ввод возраста и пола
	fmt.Scan(&age, &gender)

	// Проверяем условия: возраст от 12 до 18 включительно и пол "m"
	if age >= 12 && age <= 18 && gender == "m" {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

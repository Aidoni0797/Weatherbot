package main

import "fmt"

func main() {
	fmt.Println("Hello! What is your name?")

	var name string
	fmt.Scan(&name) // пользователь вводит свое имя. Сохраняем в переменную "name"

	fmt.Println("Hello, " + name) // приветствуем пользователя. Вместо переменной подставиться его значение, то есть то что ввел пользователь.
}

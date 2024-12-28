package main

import "fmt"

func main() {
	var action string
	fmt.Scan(&action)
	if action == "Прыгать" {
		fmt.Print("Мне нравится прыгать.")
	} else {
		if action == "Плавать" {
			fmt.Print("Я люблю плавать.")
		} else {
			if action == "Летать" {
				fmt.Print("Хотел бы я научиться летать.")
			}
		}
	}
}

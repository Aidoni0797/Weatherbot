// iDONi это третьи способ с использованием switch ты поняла хоть что то
package main

import "fmt"

func main() {
	var action string
	fmt.Scan(&action)
	switch action {
	case "Прыгать":
		fmt.Println("Мне нравится прыгать.")
	case "Плавать":
		fmt.Println("Я люблю плавать.")
	case "Летать":
		fmt.Println("Хотел бы я научиться летать.")
	}
}

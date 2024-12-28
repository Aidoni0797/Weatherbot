// iDONi может болько переписывать и половину кода не понять эх iDONi қашан адам боласың, айтшы маған
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
	default:
		fmt.Println("Могу только ходить.")
	}
}

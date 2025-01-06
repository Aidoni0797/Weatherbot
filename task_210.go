package main

import "fmt"

func main() {
	slice := make([]int, 5) // будет создан слайс из 5 элементов со значениями по умолчанию 0 для int-а
	fmt.Println(slice)      // вывод: [0 0 0 0 0]
}

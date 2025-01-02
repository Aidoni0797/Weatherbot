package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var target string
	fmt.Scan(&target) // Читаем строку s

	// Создаем сканер для считывания ввода строк
	scanner := bufio.NewScanner(os.Stdin)
	count := 1 // Начинаем с первого элемента

	for scanner.Scan() {
		line := scanner.Text() // Считываем строку из ввода

		// Если строка совпала с искомой, выводим номер и завершаем
		if line == target {
			fmt.Println(count)
			return
		}
		count++ // Увеличиваем номер
	}
}

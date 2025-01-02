package main

import "fmt"

func main() {
	// Строка для хранения результатов
	var result string

	// Перебираем все двузначные числа
	for i := 10; i <= 99; i++ {
		// Получаем цифры числа
		a := i / 10 // Первая цифра
		b := i % 10 // Вторая цифра

		// Проверяем условие: число равно удвоенному произведению цифр
		if i == 2*a*b {
			if result != "" {
				result += ","
			}
			result += fmt.Sprint(i)
		}
	}

	// Выводим результат
	fmt.Println(result)
}

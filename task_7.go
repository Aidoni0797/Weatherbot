package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) // создаем экземпляр структуры bufio.Scanner для чтения из консоли
	fmt.Println("Введите строку с пробелами:")
	_ = scanner.Scan()     // приложение остановится, пока пользователь не введет строку и не нажмет Enter
	name := scanner.Text() // сохраняем введенную строку целиком в переменную name
	fmt.Println("Вы ввели:", name)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Считываем строку-разделитель
	separator := strings.TrimSpace(readLine(reader))

	// Считываем три строки
	str1 := strings.TrimSpace(readLine(reader))
	str2 := strings.TrimSpace(readLine(reader))
	str3 := strings.TrimSpace(readLine(reader))

	// Выводим строки через разделитель
	result := strings.Join([]string{str1, str2, str3}, separator)
	fmt.Println(result)
}

// Функция для считывания строки
func readLine(reader *bufio.Reader) string {
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

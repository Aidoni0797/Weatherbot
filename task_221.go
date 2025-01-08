package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// для корректного ввода строки с пробелами воспользуемся пакетом bufio
	myscanner := bufio.NewScanner(os.Stdin)
	myscanner.Scan()
	input := myscanner.Text()

	stringNumbers := strings.Split(input, " ") // функция Split из пакета strings, разбивает строку из переменной str на коллекцию строк. Разделителем будет пробел
	var numbers []int
	for _, stringNumber := range stringNumbers {
		number, _ := strconv.Atoi(stringNumber) // преобразуем строку в число
		numbers = append(numbers, number)       // добавляем значения в коллекцию
	}
	length := len(numbers)
	for i := 0; i < length/2; i++ {
		tmp := numbers[i]
		numbers[i] = numbers[length-1-i]
		numbers[length-1-i] = tmp
	}
	for _, number := range numbers {
		fmt.Print(number, " ")
	}
}

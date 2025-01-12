package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// для коректного ввода строки с пробелами воспользуемя пакетом bufio
	myscanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите через пробел кол-во строк и столцов: ")
	myscanner.Scan()
	input := myscanner.Text()

	stringNumbers := strings.Split(input, " ")        // функция Split из пакета strings, разбивает строку из переменной input на коллекцию строк. Резделителем будет пробел
	rowsCounts, _ := strconv.Atoi(stringNumbers[0])   // преобразуем первый (индекс 0) элемент коллекции строк, сохраняем в переменную в которой будет хранится кол-во строк
	columnsCount, _ := strconv.Atoi(stringNumbers[1]) // преобразуем второй (индекс 1) элемент коллекции строк, сохраняем в переменную в которой будет хранится кол-во столбцов

	a := make([][]int, rowsCounts)
	fmt.Println("i j")
	for i := 0; i < rowsCounts; i++ {
		a[i] = make([]int, columnsCount)
		for j := 0; j < columnsCount; j++ {
			fmt.Print("[", i, ",", j, "] = ")
			fmt.Scan(&a[i][j])
		}
	}
	fmt.Println(a)
}

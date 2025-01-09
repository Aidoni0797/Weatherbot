package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Чтение входных данных
	reader := bufio.NewReader(os.Stdin)
	nInput, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nInput))

	arrayInput, _ := reader.ReadString('\n')
	arrayStr := strings.Fields(arrayInput)

	// Преобразование строкового массива в числовой
	array := make([]int, n)
	for i := 0; i < n; i++ {
		array[i], _ = strconv.Atoi(arrayStr[i])
	}

	// Подсчёт частоты каждого числа
	counts := make(map[int]int)
	for _, num := range array {
		counts[num]++
	}

	// Вывод элементов, которые встречаются только один раз
	var result []string
	for _, num := range array {
		if counts[num] == 1 {
			result = append(result, strconv.Itoa(num))
		}
	}

	// Печать результата
	fmt.Print(strings.Join(result, " "))
}

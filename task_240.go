package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	// Считываем массив
	matrix := make([][]int, n)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		scanner.Scan()
		row := strings.Fields(scanner.Text())
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			matrix[i][j], _ = strconv.Atoi(row[j])
		}
	}

	// Проверка симметричности относительно главной диагонали
	isSymmetric := true
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ { // Проверяем элементы ниже диагонали
			if matrix[i][j] != matrix[j][i] {
				isSymmetric = false
				break
			}
		}
		if !isSymmetric {
			break
		}
	}

	// Вывод результата
	if isSymmetric {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

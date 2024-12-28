package main

import (
	"fmt"
)

func main() {
	var bishopRow, bishopCol, pieceRow, pieceCol int

	// Ввод координат слона

	fmt.Scan(&bishopRow, &bishopCol)

	// Ввод координат другой фигуры

	fmt.Scan(&pieceRow, &pieceCol)

	// Проверка, может ли слон побить фигуру
	if abs(bishopRow-pieceRow) == abs(bishopCol-pieceCol) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

// Функция для вычисления модуля числа
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

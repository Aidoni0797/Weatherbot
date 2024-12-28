package main

import (
	"fmt"
)

func main() {
	var rookRow, rookCol, pieceRow, pieceCol int

	// Ввод координат ладьи

	fmt.Scan(&rookRow, &rookCol)

	// Ввод координат фигуры

	fmt.Scan(&pieceRow, &pieceCol)

	// Проверка, может ли ладья побить фигуру
	if rookRow == pieceRow || rookCol == pieceCol {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

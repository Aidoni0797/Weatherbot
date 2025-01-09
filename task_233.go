package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Считываем число N
	line, _ := reader.ReadString('\n')
	N, _ := strconv.Atoi(strings.TrimSpace(line))

	// Считываем массив чисел
	line, _ = reader.ReadString('\n')
	elements := strings.Fields(line)

	numbers := make(map[int]bool)
	for i := 0; i < N; i++ {
		num, _ := strconv.Atoi(elements[i])
		if numbers[num] {
			fmt.Print("YES")
			return
		}
		numbers[num] = true
	}

	fmt.Print("NO")
}

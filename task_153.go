// если бы не chatgpt, то система не пропустила бы, какой ты молодец chtgpt, я тобою восхищаюсь
package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	power := 1
	for power <= N {
		fmt.Println(power)
		power *= 2
	}
}

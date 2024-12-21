// Здесь _, _ = fmt.Scan(&lastName, &firstName, &patronymic) указывает,
// что мы сознательно игнорируем количество считанных слов и возможные ошибки.

// сознательное игнорирование какое прекрасно
package main

import "fmt"

func main() {
	var lastName, firstName, patronymic string
	_, _ = fmt.Scan(&lastName, &firstName, &patronymic)
}

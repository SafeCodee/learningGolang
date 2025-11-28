package formatter

import "fmt"

/*
=== ФАЙЛ: formatter/formatter.go ===

package formatter

// TODO: Импортируй "fmt"

// TODO: Создай экспортируемую функцию FormatResult(operation string, result int) string
// которая возвращает строку в формате: "{operation}: {result}"
// Пример: "Сложение: 15"
// Подсказка: используй fmt.Sprintf

// TODO: Создай приватную функцию padString(s string) string
// которая добавляет ">>> " в начало строки
// Эта функция не будет доступна извне пакета

// TODO: Создай экспортируемую функцию FormatError(err error) string
// которая возвращает строку: "Ошибка: {текст ошибки}"
*/
func FormatResult(operation string, result int) string {
	return fmt.Sprintf("%s: %d", operation, result)
}

func padString(s string) string {
	return ">>>" + s
}

func FormatError(err error) string {
	return fmt.Sprintf("Ошибка: %s", err)
}

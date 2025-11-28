package validator

import "errors"

/*
=== ФАЙЛ: internal/validator/validator.go ===

package validator

// TODO: Импортируй "errors"

// TODO: Создай экспортируемую функцию ValidatePositive(n int) error
// которая возвращает ошибку если n <= 0
// Текст ошибки: "число должно быть положительным"

// TODO: Создай приватную функцию validate(n int) bool
// которая возвращает true если n > 0
// Эта функция используется внутри ValidatePositive

Подсказка: хотя функция ValidatePositive экспортируется из пакета validator,
она доступна ТОЛЬКО внутри модуля chapter_12_calculator (из-за internal/)
*/

func ValidatePositive(n int) error {
	if !validate(n) {
		return errors.New("число должно быть положительным")
	}

	return nil
}

func validate(n int) bool {
	if n > 0 {
		return true
	}

	return false
}

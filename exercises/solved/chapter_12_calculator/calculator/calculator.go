package calculator

import "errors"

/*
=== ФАЙЛ: calculator/calculator.go ===

package calculator

// TODO: Создай экспортируемую функцию Add(a, b int) int
// которая складывает два числа

// TODO: Создай экспортируемую функцию Subtract(a, b int) int
// которая вычитает b из a

// TODO: Создай экспортируемую функцию Multiply(a, b int) int
// которая умножает два числа

// TODO: Создай экспортируемую функцию Divide(a, b int) (int, error)
// Функция должна:
// 1. Использовать приватную функцию isValidDivision(b int) error
//    для проверки деления на ноль
// 2. Если ошибка — вернуть 0 и ошибку
// 3. Иначе вернуть результат деления и nil

// TODO: Создай ПРИВАТНУЮ функцию isValidDivision(divisor int) error
// которая возвращает ошибку если divisor == 0
// Подсказка: используй errors.New("деление на ноль невозможно")
*/

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) (int, error) {
	if err := isValidDivision(b); err != nil {
		return 0, err
	}

	return a / b, nil
}

func isValidDivision(divisor int) error {
	if divisor == 0 {
		return errors.New("деление на ноль невозможно")
	}

	return nil
}

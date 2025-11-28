package main

/*
Задание: Создание многопакетного проекта "Calculator App"

Структура проекта (создай её вручную):

chapter_12_calculator/
├── go.mod                    # Создай с помощью: go mod init chapter_12_calculator
├── main.go                   # Этот файл (точка входа)
├── calculator/               # Пакет для математических операций
│   └── calculator.go
├── formatter/                # Пакет для форматирования результатов
│   └── formatter.go
└── internal/
    └── validator/            # Внутренний пакет для валидации
        └── validator.go

ВАЖНО: Скопируй этот файл целиком в директорию chapter_12_calculator/
и назови его main.go, затем создай остальные файлы по инструкции ниже.
*/

// TODO: Импортируй необходимые пакеты:
// - fmt (стандартная библиотека)
// - chapter_12_calculator/calculator (твой пакет)
// - chapter_12_calculator/formatter (твой пакет)

func main() {
	// TODO: Используй функции из пакета calculator для:
	// 1. Сложения 10 и 5
	// 2. Вычитания 10 - 3
	// 3. Умножения 4 * 7
	// 4. Деления 20 / 4

	// TODO: Используй функцию FormatResult из пакета formatter
	// для вывода результатов в формате: "Результат: X"

	// TODO: Попробуй использовать приватную функцию validate()
	// из internal/validator напрямую — убедись что получаешь ошибку компиляции
	// (закомментируй эту строку после проверки)
}

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

/*
=== ДОПОЛНИТЕЛЬНОЕ ЗАДАНИЕ (необязательно): ===

1. Добавь в calculator.go функцию Power(base, exp int) int
   которая возводит base в степень exp (используй цикл)

2. Создай новый файл calculator/stats.go в том же пакете calculator
   Добавь функцию Average(numbers []int) float64
   Проверь что две функции из разных файлов одного пакета работают вместе

3. Попробуй импортировать internal/validator в main.go
   Убедись что компилятор запрещает это (потом удали импорт)
*/

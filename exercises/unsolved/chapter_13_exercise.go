package main

import (
	"fmt"
	"os"
	"time"
)

// Задание 1: Defer и порядок выполнения
// TODO: Создай функцию deferOrder(), которая демонстрирует порядок выполнения defer (LIFO)
// Функция должна вызвать defer 5 раз с разными сообщениями и показать порядок их выполнения

// Задание 2: Defer с ресурсами
// TODO: Создай функцию writeToFile(filename string, content string) error
// Функция должна:
// - Создать файл (os.Create)
// - Использовать defer для гарантированного закрытия файла
// - Записать content в файл
// - Вернуть ошибку, если что-то пошло не так

// Задание 3: Defer и изменение именованного возвращаемого значения
// TODO: Создай функцию modifyReturn() (result int)
// Функция должна:
// - Установить result = 10
// - В defer изменить result на 20
// - Вернуть result (должно вернуться 20, а не 10)

// Задание 4: Измерение времени выполнения
// TODO: Создай функцию measureExecutionTime(name string) func()
// Функция должна:
// - Запомнить время старта (time.Now())
// - Вернуть анонимную функцию, которая выводит elapsed time
// - Эта функция будет использоваться с defer для измерения времени

// TODO: Создай функцию slowOperation(), которая:
// - Использует defer measureExecutionTime("slowOperation")()
// - Делает time.Sleep(100 * time.Millisecond)

// Задание 5: Безопасное деление с panic
// TODO: Создай функцию divide(a, b int) int
// Функция должна:
// - Проверить если b == 0, вызвать panic("division by zero")
// - Иначе вернуть a / b

// Задание 6: Безопасное деление с recover
// TODO: Создай функцию safeDivide(a, b int) (result int, err error)
// Функция должна:
// - Использовать defer с recover для перехвата паники
// - Если паника произошла, установить err = fmt.Errorf("ошибка: %v", r)
// - Вызвать внутри divide(a, b)
// - Вернуть результат или ошибку

// Задание 7: Доступ к элементу слайса с recover
// TODO: Создай функцию safeGetElement(slice []int, index int) (value int, err error)
// Функция должна:
// - Использовать defer с recover для перехвата паники при выходе за границы
// - Попытаться получить slice[index]
// - Если паника (index out of range), вернуть ошибку
// - Если всё ок, вернуть значение

// Задание 8: Defer в цикле (правильное использование)
// TODO: Создай функцию processFiles(filenames []string) error
// Функция должна:
// - Для каждого файла вызвать отдельную функцию processFile(filename string) error
// - processFile должна открыть файл, использовать defer для закрытия, и прочитать первую строку
// - Если любой файл вызвал ошибку, вернуть её

// TODO: Создай вспомогательную функцию processFile(filename string) error

// Задание 9: Логирование выхода из функции
// TODO: Создай функцию calculateWithLog(a, b int) (result int, err error)
// Функция должна:
// - В defer логировать результат выполнения функции (успех или ошибка)
// - Если b == 0, вернуть ошибку
// - Иначе вернуть a / b

func main() {
	fmt.Println("=== Задание 1: Порядок выполнения defer ===")
	// TODO: Вызови deferOrder()

	fmt.Println("\n=== Задание 2: Defer с ресурсами ===")
	// TODO: Вызови writeToFile("test.txt", "Hello from defer!")
	// TODO: Проверь ошибку и выведи результат

	fmt.Println("\n=== Задание 3: Defer и изменение возвращаемого значения ===")
	// TODO: Вызови modifyReturn() и выведи результат

	fmt.Println("\n=== Задание 4: Измерение времени выполнения ===")
	// TODO: Вызови slowOperation()

	fmt.Println("\n=== Задание 5: Panic при делении на ноль ===")
	// TODO: Вызови divide(10, 2)
	// TODO: Раскомментируй следующую строку чтобы увидеть панику:
	// divide(10, 0)

	fmt.Println("\n=== Задание 6: Безопасное деление с recover ===")
	// TODO: Вызови safeDivide(10, 2) и выведи результат
	// TODO: Вызови safeDivide(10, 0) и выведи ошибку

	fmt.Println("\n=== Задание 7: Безопасный доступ к слайсу ===")
	// TODO: Создай слайс []int{1, 2, 3, 4, 5}
	// TODO: Вызови safeGetElement(slice, 2) - должно вернуть 3
	// TODO: Вызови safeGetElement(slice, 10) - должно вернуть ошибку

	fmt.Println("\n=== Задание 8: Defer в цикле ===")
	// TODO: Создай несколько тестовых файлов или используй существующие
	// TODO: Вызови processFiles([]string{"test.txt", "go.mod"})

	fmt.Println("\n=== Задание 9: Логирование выхода из функции ===")
	// TODO: Вызови calculateWithLog(10, 2)
	// TODO: Вызови calculateWithLog(10, 0)

	// Очистка: удаляем тестовый файл
	os.Remove("test.txt")
}

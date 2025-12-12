package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"time"
)

// Задание 1: Defer и порядок выполнения
// TODO: Создай функцию deferOrder(), которая демонстрирует порядок выполнения defer (LIFO)
// Функция должна вызвать defer 5 раз с разными сообщениями и показать порядок их выполнения
func deferOrder() {
	fmt.Println("main execution from deferOrder")

	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")
	defer fmt.Println("defer 4")
	defer func() {
		fmt.Println("defer 5")
	}()
}

// Задание 2: Defer с ресурсами
// TODO: Создай функцию writeToFile(filename string, content string) error
// Функция должна:
// - Создать файл (os.Create)
// - Использовать defer для гарантированного закрытия файла
// - Записать content в файл
// - Вернуть ошибку, если что-то пошло не так
func writeToFile(filename string, content string) error {
	file, err := os.Create(filename)
	defer file.Close()

	if err != nil {
		return err
	}

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

// Задание 3: Defer и изменение именованного возвращаемого значения
// TODO: Создай функцию modifyReturn() (result int)
// Функция должна:
// - Установить result = 10
// - В defer изменить result на 20
// - Вернуть result (должно вернуться 20, а не 10)

func modifyReturn() (result int) {
	result = 10
	defer func() {
		result = 20
	}()

	return result
}

// Задание 4: Измерение времени выполнения
// TODO: Создай функцию measureExecutionTime(name string) func()
// Функция должна:
// - Запомнить время старта (time.Now())
// - Вернуть анонимную функцию, которая выводит elapsed time
// - Эта функция будет использоваться с defer для измерения времени
func measureExecutionTime(name string) func() {
	startTime := time.Now()

	return func() {
		fmt.Printf("Elapsed %d", time.Since(startTime))
	}
}

// TODO: Создай функцию slowOperation(), которая:
// - Использует defer measureExecutionTime("slowOperation")()
// - Делает time.Sleep(100 * time.Millisecond)
func slowOperation() {

	defer measureExecutionTime("slowOperation")()
	time.Sleep(100 * time.Millisecond)
}

// Задание 5: Безопасное деление с panic
// TODO: Создай функцию divide(a, b int) int
// Функция должна:
// - Проверить если b == 0, вызвать panic("division by zero")
// - Иначе вернуть a / b
func divide2(a, b int) int {
	if b == 0 {
		panic("division by zero")
	}

	return a / b
}

// Задание 6: Безопасное деление с recover
// TODO: Создай функцию safeDivide(a, b int) (result int, err error)
// Функция должна:
// - Использовать defer с recover для перехвата паники
// - Если паника произошла, установить err = fmt.Errorf("ошибка: %v", r)
// - Вызвать внутри divide(a, b)
// - Вернуть результат или ошибку
func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("ошибка: %v", r)
		}

	}()

	return divide2(a, b), nil
}

// Задание 7: Доступ к элементу слайса с recover
// TODO: Создай функцию safeGetElement(slice []int, index int) (value int, err error)
// Функция должна:
// - Использовать defer с recover для перехвата паники при выходе за границы
// - Попытаться получить slice[index]
// - Если паника (index out of range), вернуть ошибку
// - Если всё ок, вернуть значение
func safeGetElement(slice []int, index int) (value int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("index out of range")
		}
	}()

	return slice[index], nil
}

// Задание 8: Defer в цикле (правильное использование)
// TODO: Создай функцию processFiles(filenames []string) error
// Функция должна:
// - Для каждого файла вызвать отдельную функцию processFile(filename string) error
// - processFile должна открыть файл, использовать defer для закрытия, и прочитать первую строку
// - Если любой файл вызвал ошибку, вернуть её
func processFiles(filenames []string) error {
	for _, fileName := range filenames {
		err := processFile(fileName)
		if err != nil {
			return err
		}
	}

	return nil
}

func processFile(filename string) error {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return err
	}

	// Create a new reader
	reader := bufio.NewReader(file)

	var line string
	for {
		// Read until we encounter a newline character
		line, err = reader.ReadString('n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
		}

		// Print the line
		break
	}

	fmt.Println(line)
	return nil
}

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
	deferOrder()

	fmt.Println("\n=== Задание 2: Defer с ресурсами ===")
	// TODO: Вызови writeToFile("test.txt", "Hello from defer!")
	// TODO: Проверь ошибку и выведи результат
	err := writeToFile("test.txt", "Hello from defer!")
	if err != nil {
		fmt.Printf("There was an error %s\n", err)
	} else {
		fmt.Printf("There was not an error \n")
	}

	fmt.Println("\n=== Задание 3: Defer и изменение возвращаемого значения ===")
	// TODO: Вызови modifyReturn() и выведи результат
	fmt.Printf("modifyReturn %d", modifyReturn())

	fmt.Println("\n=== Задание 4: Измерение времени выполнения ===")
	// TODO: Вызови slowOperation()
	slowOperation()

	fmt.Println("\n=== Задание 5: Panic при делении на ноль ===")
	// TODO: Вызови divide(10, 2)
	// TODO: Раскомментируй следующую строку чтобы увидеть панику:
	// divide(10, 0)
	//divide2(10, 0)
	fmt.Println("\n=== Задание 6: Безопасное деление с recover ===")
	// TODO: Вызови safeDivide(10, 2) и выведи результат
	// TODO: Вызови safeDivide(10, 0) и выведи ошибку
	fmt.Println(safeDivide(10, 2))
	fmt.Println(safeDivide(10, 0))
	fmt.Println("\n=== Задание 7: Безопасный доступ к слайсу ===")
	// TODO: Создай слайс []int{1, 2, 3, 4, 5}
	// TODO: Вызови safeGetElement(slice, 2) - должно вернуть 3
	// TODO: Вызови safeGetElement(slice, 10) - должно вернуть ошибку
	var slice []int = []int{1, 2, 3, 4, 5}
	fmt.Println(safeGetElement(slice, 2))
	fmt.Println(safeGetElement(slice, 10))

	fmt.Println("\n=== Задание 8: Defer в цикле ===")
	// TODO: Создай несколько тестовых файлов или используй существующие
	// TODO: Вызови processFiles([]string{"test.txt", "go.mod"})
	err = processFiles([]string{"test1.txt", "test2.txt"})
	if err != nil {
		fmt.Printf("error during processFiles %s\n", err)
	}

	err = processFiles([]string{"test1.txt", "test123.txt"})
	if err != nil {
		fmt.Printf("error during processFiles %s\n", err)
	}
	fmt.Println("\n=== Задание 9: Логирование выхода из функции ===")
	// TODO: Вызови calculateWithLog(10, 2)
	// TODO: Вызови calculateWithLog(10, 0)

	// Очистка: удаляем тестовый файл
	os.Remove("test.txt")
}

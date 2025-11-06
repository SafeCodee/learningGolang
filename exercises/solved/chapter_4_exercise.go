package main

import "fmt"

// TODO: Напиши функцию average, которая принимает вариадический параметр numbers типа float64
// и возвращает среднее значение (float64) и количество чисел (int).
// Если чисел нет, верни (0.0, 0)

func average(numbers ...float64) (float64, int) {
	length := len(numbers)
	if numbers == nil || length == 0 {
		return 0.0, 0
	}

	sum := 0.0
	for _, value := range numbers {
		sum += value
	}

	return sum / float64(length), length

}

// TODO: Напиши функцию calculator, которая принимает два числа (a, b int) и операцию (operation string).
// Операция может быть: "+", "-", "*", "/"
// Функция должна возвращать результат (int) и ошибку (error).
// Если операция "/" и b == 0, верни ошибку "division by zero"
// Если операция неизвестна, верни ошибку "unknown operation: <operation>"
// Подсказка: используй fmt.Errorf для создания ошибки
func calculator(a, b int, operation string) (int, error) {
	switch operation {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		{
			if b == 0 {
				return 0, fmt.Errorf("division by zero")
			}

			return a / b, nil
		}
	default:
		return 0, fmt.Errorf("unknown operation: %s", operation)
	}
}

// TODO: Напиши функцию makeMultiplier, которая принимает factor (int)
// и возвращает функцию, умножающую число на factor.
// Пример: double := makeMultiplier(2)
//
//	fmt.Println(double(5)) // 10
//
// Это замыкание (closure)!
func makeMultiplier(factor int) func(int) int {
	return func(value int) int {
		return value * factor
	}
}

// TODO: Напиши функцию processNumbers, которая принимает слайс чисел []int
// и функцию-обработчик func(int) int, применяет обработчик к каждому элементу
// и возвращает новый слайс с результатами.
// Это пример функции высшего порядка (higher-order function)

func processNumbers(numbers []int, function func(value int) int) []int {
	var result []int

	for _, value := range numbers {
		result = append(result, function(value))
	}

	return result
}

func averageReport(array ...float64) {
	avg, length := average(array...)
	fmt.Printf("Average of %v is %.2f (count: %d)\n", array, avg, length)
}

func main() {

	defer fmt.Println("Program finished!")
	// TODO: Протестируй функцию average
	// Вызови с разными наборами чисел: (10, 20, 30), (5.5), ()
	// Выведи результат в формате: "Average of [10 20 30] is 20.00 (count: 3)"
	var array = []float64{10, 20, 30}
	averageReport(array...)

	array = []float64{5.5}
	averageReport(array...)

	array = []float64{}
	averageReport(array...)

	// TODO: Протестируй функцию calculator
	// Попробуй разные операции: +, -, *, /
	// Протестируй деление на ноль
	// Протестируй неизвестную операцию
	// Обрабатывай ошибки! if err != nil { ... }
	a := 12
	b := 3
	calculatorReport(a, b, "+", "%d + %d = %d\n")
	calculatorReport(a, b, "-", "%d - %d = %d\n")
	calculatorReport(a, b, "*", "%d * %d = %d\n")
	calculatorReport(a, b, "/", "%d / %d = %d\n")
	calculatorReport(a, 0, "/", "%d / %d = %d\n")
	calculatorReport(a, b, "!", "%d ! %d = %d\n")

	// TODO: Протестируй функцию makeMultiplier
	// Создай функции: double (множитель 2), triple (множитель 3)
	// Примени их к числу 7 и выведи результаты
	double := makeMultiplier(2)
	fmt.Printf("double %d\n", double(7))
	triple := makeMultiplier(3)
	fmt.Printf("triple %d\n", triple(7))

	// TODO: Протестируй функцию processNumbers
	// Создай слайс чисел: []int{1, 2, 3, 4, 5}
	// Создай функцию-обработчик, которая возводит число в квадрат
	// Передай в processNumbers и выведи результат
	ints := []int{1, 2, 3, 4, 5}
	numbers := processNumbers(ints, func(value int) int {
		return value * value
	})
	fmt.Println(numbers)

	// BONUS TODO: Используй defer для вывода "Program finished!" в самом конце main
	// defer должен быть объявлен В НАЧАЛЕ функции main, но выполнится в конце

}

func calculatorReport(a int, b int, operation string, formating string) {
	result, err := calculator(a, b, operation)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf(formating, a, b, result)
	}
}

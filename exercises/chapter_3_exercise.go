package main

import "fmt"

func main() {
	// TODO: Создай переменную score типа int со значением 85
	score := 85
	// TODO: Используй if-else if-else для определения оценки:
	// score >= 90: "Отлично"
	// score >= 75: "Хорошо"
	// score >= 60: "Удовлетворительно"
	// иначе: "Неудовлетворительно"
	if score >= 90 {
		fmt.Println("Отлично")
	} else if score >= 75 {
		fmt.Println("Хорошо")
	} else if score >= 60 {
		fmt.Println("Удовлетворительно")
	} else {
		fmt.Println("Неудовлетворительно")
	}
	// TODO: Используй классический цикл for для вывода чисел от 1 до 10
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	// TODO: Создай слайс чисел: numbers := []int{5, 10, 15, 20, 25}
	// Используй for-range для вывода индекса и значения каждого элемента
	// Формат: "Индекс 0: значение 5"
	numbers := []int{5, 10, 15, 20, 25}
	for index, value := range numbers {
		fmt.Printf("Индекс %d: значение %d\n", index, value)
	}
	// TODO: Создай переменную sum и используй for-range для подсчёта суммы всех элементов numbers
	// Подсказка: игнорируй индекс используя _
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	// TODO: Используй цикл for как while:
	// Создай переменную counter = 1
	// Пока counter <= 5, выводи counter и удваивай его (counter *= 2)
	// Ожидаемый вывод: 1, 2, 4
	counter := 1
	for counter <= 5 {
		fmt.Printf("%d ", counter)
		counter *= 2
	}
	fmt.Println()
	// TODO: Создай переменную dayOfWeek со значением "Friday"
	// Используй switch для вывода:
	// Monday-Friday: "Рабочий день"
	// Saturday, Sunday: "Выходной"
	friday := "Friday"
	switch friday {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Рабочий день")
	case "Saturday", "Sunday":
		fmt.Println("Выходной")
	}
	// TODO: Создай switch БЕЗ выражения для проверки score:
	// >= 90: "A"
	// >= 80: "B"
	// >= 70: "C"
	// >= 60: "D"
	// иначе: "F"
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 70:
		fmt.Println("С")
	case score >= 60:
		fmt.Println("В")
	default:
		fmt.Println("F")
	}

	// БОНУС: Используй вложенные циклы с меткой
	// Создай таблицу умножения 3x3, но прерви выполнение когда i*j >= 6
	// Используй метку outer и break outer
outer:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i*j >= 6 {
				break outer
			}
			fmt.Printf("%d ", i*j)
		}
		fmt.Println()
	}
}

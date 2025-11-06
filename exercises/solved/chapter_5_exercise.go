package main

import "fmt"

func main() {
	// Часть 1: Массивы
	// TODO: Создай массив из 5 строк с названиями дней недели (Пн, Вт, Ср, Чт, Пт)
	// Используй инициализацию с [...]
	days := [...]string{"Пн", "Вт", "Ср", "Чт", "Пт"}
	// TODO: Выведи третий элемент массива (среду)
	fmt.Println(days[2])
	// Часть 2: Создание и работа со слайсами
	// TODO: Создай слайс целых чисел от 1 до 10 используя литерал слайса []int{...}
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// TODO: Выведи длину (len) и ёмкость (cap) слайса
	fmt.Printf("length=%d, capacity=%d\n", len(numbers), cap(numbers))
	// TODO: Используй цикл range для вывода каждого элемента в формате "Индекс: значение"
	for i, value := range numbers {
		fmt.Printf("%d: %d\n", i, value)
	}
	// Часть 3: Append и динамическое изменение
	// TODO: Создай пустой слайс строк
	var strings []string
	// TODO: Добавь в него три города: "Москва", "Питер", "Казань"
	// Используй append несколько раз
	strings = append(strings, "Москва")
	strings = append(strings, "Питер")
	strings = append(strings, "Казань")

	// TODO: Выведи получившийся слайс
	fmt.Println(strings)
	// Часть 4: Slicing (получение подслайса)
	// TODO: Создай слайс чисел от 0 до 9: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	numbersSlice := make([]int, 10)
	for i := 0; i < 10; i++ {
		numbersSlice[i] = i
	}
	// TODO: Создай подслайс с элементами от индекса 3 до 7 (не включая 7)
	subSlice := numbersSlice[3:7]
	// TODO: Выведи подслайс
	fmt.Println(subSlice)
	// TODO: Измени первый элемент подслайса на 999
	subSlice[0] = 999
	// TODO: Выведи оригинальный слайс — он тоже должен измениться!
	fmt.Println(numbersSlice)
	// Часть 5: Copy (независимое копирование)
	// TODO: Создай слайс original := []int{1, 2, 3, 4, 5}
	original := []int{1, 2, 3, 4, 5}
	// TODO: Создай новый слайс copied такой же длины с помощью make
	copied := make([]int, len(original))
	// TODO: Скопируй элементы из original в copied используя функцию copy
	copy(copied, original)
	// TODO: Измени первый элемент copied на 100
	copied[0] = 100
	// TODO: Выведи оба слайса — original не должен измениться
	fmt.Println(original)
	fmt.Println(copied)
	// Часть 6: Практическая задача
	// TODO: Напиши функцию filterEven, которая принимает слайс int
	// и возвращает новый слайс только с чётными числами
	// Подсказка: используй append для добавления в результирующий слайс

	// TODO: Протестируй функцию на слайсе []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// Должен вернуться [2, 4, 6, 8, 10]
	fmt.Println(filterEven([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}

// TODO: Напиши функцию filterEven здесь
func filterEven(numbers []int) []int {
	var result []int
	for _, value := range numbers {
		if value%2 == 0 {
			result = append(result, value)
		}
	}

	return result
}

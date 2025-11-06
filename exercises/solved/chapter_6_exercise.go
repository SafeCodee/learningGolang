package main

import (
	"fmt"
	"sort"
	"strings"
)

// TODO: Создай функцию countWords, которая принимает строку текста
// и возвращает мапу где ключ - слово, значение - количество его появлений
// Например: "hello world hello" -> map["hello":2, "world":1]
// Подсказка: используй strings.Fields() для разделения строки на слова
func countWords(text string) map[string]int {
	result := make(map[string]int)
	words := strings.Fields(text)
	for _, word := range words {
		result[word] += 1
	}

	return result
}

// TODO: Создай функцию mergeMaps, которая принимает две мапы map[string]int
// и возвращает новую мапу с объединёнными данными
// Если ключ есть в обеих мапах - суммируй значения
// Например: map["a":1, "b":2] + map["b":3, "c":4] -> map["a":1, "b":5, "c":4]
func mergeMaps(m1, m2 map[string]int) map[string]int {
	result := make(map[string]int)
	for key, value := range m1 {
		result[key] = value
	}
	for key, value := range m2 {
		result[key] += value
	}

	return result
}

// TODO: Создай функцию invertMap, которая "переворачивает" мапу:
// ключи становятся значениями, значения становятся ключами
// Например: map["a":1, "b":2] -> map[1:"a", 2:"b"]
// ⚠️ Что если несколько ключей имеют одинаковое значение?
// Решение: последний встреченный ключ побеждает (или можешь сохранить любой)
func invertMap(m map[string]int) map[int]string {
	result := make(map[int]string)

	for key, value := range m {
		result[value] = key
	}

	return result
}

// TODO: Создай функцию getTopScores, которая принимает мапу имя->балл
// и число n, возвращает слайс из n имён с наивысшими баллами
// Например: map["Alice":90, "Bob":85, "Charlie":95], n=2 -> ["Charlie", "Alice"]
// Подсказка: нужно будет отсортировать (import "sort")
func getTopScores(scores map[string]int, n int) []string {
	// Создаём структуру для пары имя-балл
	type student struct {
		name  string
		score int
	}

	// Собираем всех студентов в слайс
	students := make([]student, 0, len(scores))
	for name, score := range scores {
		students = append(students, student{name, score})
	}

	// Сортируем по баллу (по убыванию)
	sort.Slice(students, func(i, j int) bool {
		return students[i].score > students[j].score
	})

	// Берём первые n имён
	result := make([]string, 0, n)
	for i := 0; i < n && i < len(students); i++ {
		result = append(result, students[i].name)
	}

	return result
}

func main() {
	fmt.Println("=== Задание 1: Подсчёт слов ===")
	text := "go is awesome go go is simple"
	wordCount := countWords(text)
	fmt.Printf("Текст: %s\n", text)
	fmt.Printf("Результат: %v\n\n", wordCount)
	// Ожидается: map[go:3 is:2 awesome:1 simple:1]

	fmt.Println("=== Задание 2: Объединение мап ===")
	map1 := map[string]int{"a": 1, "b": 2, "c": 3}
	map2 := map[string]int{"b": 5, "c": 7, "d": 9}
	merged := mergeMaps(map1, map2)
	fmt.Printf("Map1: %v\n", map1)
	fmt.Printf("Map2: %v\n", map2)
	fmt.Printf("Merged: %v\n\n", merged)
	// Ожидается: map[a:1 b:7 c:10 d:9]

	fmt.Println("=== Задание 3: Инверсия мапы ===")
	original := map[string]int{"Alice": 25, "Bob": 30, "Charlie": 25}
	inverted := invertMap(original)
	fmt.Printf("Original: %v\n", original)
	fmt.Printf("Inverted: %v\n\n", inverted)
	// Ожидается: map[25:Charlie 30:Bob] или map[25:Alice 30:Bob]
	// (порядок случайный, может быть Alice или Charlie для ключа 25)

	fmt.Println("=== Задание 4: Топ-N баллов ===")
	scores := map[string]int{
		"Alice":   90,
		"Bob":     85,
		"Charlie": 95,
		"David":   88,
		"Eve":     92,
	}
	top3 := getTopScores(scores, 3)
	fmt.Printf("Scores: %v\n", scores)
	fmt.Printf("Top 3: %v\n", top3)
	// Ожидается: [Charlie Eve Alice] (95, 92, 90)

	fmt.Println("\n=== Дополнительная практика ===")
	// TODO: Создай мапу студентов и их оценок (как в примере выше)
	// TODO: Добавь нового студента с помощью []
	// TODO: Проверь существует ли студент "Zara" с помощью value, ok := map[key]
	// TODO: Удали одного студента с помощью delete()
	// TODO: Выведи всех студентов и их оценки используя range
	// TODO: Выведи средний балл всех студентов
	studentsMap := map[string]int{
		"Pit":   100,
		"Alice": 90,
		"Zara":  99,
	}
	studentsMap["Vova"] = 96
	_, ok := studentsMap["Zara"]
	if ok {
		fmt.Printf("Zara is our student\n")
	}

	delete(studentsMap, "Pit")
	sumScore := 0
	for name, score := range studentsMap {
		fmt.Printf("Name = %s, score = %d\n", name, score)
		sumScore += score
	}

	fmt.Printf("Avarage score %d\n", sumScore/len(studentsMap))

}

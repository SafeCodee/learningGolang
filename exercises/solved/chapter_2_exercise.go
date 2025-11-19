package main

import (
	"fmt"
	"strconv"
)

func main() {
	// TODO: Объяви переменную age типа int со значением 33 используя полную форму var
	var age int = 33
	// TODO: Объяви переменную salary типа float64 со значением 75000.50 используя короткую форму :=
	salary := 75000.50
	// TODO: Объяви переменную isEmployed типа bool со значением true
	isEmployed := true
	// TODO: Выведи все три переменные используя fmt.Printf
	// Подсказка: %d для int, %f для float64, %t для bool
	fmt.Printf("%d, %f, %t\n", age, salary, isEmployed)
	// TODO: Создай константу TaxRate со значением 0.13 (13% налог)
	const TaxRate = 0.13
	// TODO: Вычисли налог: tax = salary * TaxRate
	// Обрати внимание на типы! TaxRate - это float64, salary - тоже float64
	tax := salary * TaxRate
	// TODO: Вычисли зарплату после налогов: netSalary = salary - tax
	netSalary := salary - tax
	// TODO: Выведи результат: "Зарплата до налогов: X, налог: Y, после налогов: Z"
	// Подсказка: используй %.2f для вывода с 2 знаками после запятой
	fmt.Printf("Зарплата до налогов: %.2f, налог: %.2f, после налогов: %.2f\n", salary, tax, netSalary)

	// TODO: Попробуй создать переменную с нулевым значением (без инициализации)
	// и выведи её чтобы увидеть zero value
	var a int
	fmt.Println(a)
	// БОНУС: Попробуй сложить int и float64 напрямую и посмотри на ошибку компиляции
	// Потом исправь используя явное приведение типов
	//var result = age + salary // Раскомментируй и посмотри на ошибку
	//var result = float64(age) + salary  // Правильный вариант

	// === Работа со строками ===

	// TODO: Создай переменные firstName и lastName со значениями "Иван" и "Петров"
	firstName, lastName := "Иван", "Петров"
	// TODO: Создай fullName объединив firstName + " " + lastName
	fullName := firstName + " " + lastName
	// TODO: Выведи fullName и его длину используя len()
	// Обрати внимание: len() вернёт количество байт, не символов!
	fmt.Printf("FullName = %s, len = %d\n", fullName, len(fullName))
	// TODO: Получи первый байт строки firstName (firstName[0]) и выведи его
	// Это будет число (byte), не символ!
	fmt.Printf("firstName[0] = %d\n", firstName[0])
	// === Работа с символами (rune и byte) ===

	// TODO: Создай переменную letterA с символом 'A' (это будет rune)
	letterA := 'A'
	// TODO: Создай переменную letterRu с символом 'Я' (кириллица - тоже rune)
	letterRu := 'Я'
	// TODO: Выведи оба символа используя %c (для символа) и %d (для числового значения)
	// Например: fmt.Printf("Символ: %c, код: %d\n", letterA, letterA)
	fmt.Printf("Символ: %c, код: %d\n", letterA, letterA)
	fmt.Printf("Символ: %c, код: %d\n", letterRu, letterRu)
	// === Преобразование string ↔ int ===
	// Не забудь добавить import "strconv" в начало файла!

	// TODO: Преобразуй age (int) в строку используя strconv.Itoa()
	strAge := strconv.Itoa(age)
	// TODO: Создай строку yearStr со значением "2025"
	yearStr := "2025"
	// TODO: Преобразуй yearStr в int используя strconv.Atoi()
	// Это вернёт два значения: число и ошибку
	// Используй year, _ := strconv.Atoi(yearStr) чтобы игнорировать ошибку
	intYear, _ := strconv.Atoi(yearStr)
	// TODO: Выведи результаты конвертации
	fmt.Println(strAge, intYear)
}

package main

import "fmt"

const (
	TaxRate float64 = 0.13
)

func main() {
	// TODO: Объяви переменную age типа int со значением 33 используя полную форму var
	var age int = 33
	// TODO: Объяви переменную salary типа float64 со значением 75000.50 используя короткую форму :=
	salary := 75000.00
	// TODO: Объяви переменную isEmployed типа bool со значением true
	isEmployed := true
	// TODO: Выведи все три переменные используя fmt.Printf
	// Подсказка: %d для int, %f для float64, %t для bool
	fmt.Printf("age=%d, salary=%.2f, isEmployed=%t\n", age, salary, isEmployed)
	// TODO: Создай константу TaxRate со значением 0.13 (13% налог)

	// TODO: Вычисли налог: tax = salary * TaxRate
	// Обрати внимание на типы! TaxRate - это float64, salary - тоже float64
	tax := TaxRate * salary
	// TODO: Вычисли зарплату после налогов: netSalary = salary - tax
	netSalary := salary - tax
	// TODO: Выведи результат: "Зарплата до налогов: X, налог: Y, после налогов: Z"
	// Подсказка: используй %.2f для вывода с 2 знаками после запятой
	fmt.Printf("Зарплата до налогов: %.2f, налог: %.2f, после налогов: %.2f\n", salary, tax, netSalary)
	// TODO: Попробуй создать переменную с нулевым значением (без инициализации)
	// и выведи её чтобы увидеть zero value
	var zeroInt int
	var zeroString string
	var zeroBool bool
	fmt.Printf("Zero values: int=%d, string='%s', bool=%t\n", zeroInt, zeroString, zeroBool)

	// БОНУС: Попробуй сложить int и float64 напрямую и посмотри на ошибку компиляции
	// Потом исправь используя явное приведение типов
	//var result = age + salary // Раскомментируй и посмотри на ошибку
	//var result = float64(age) + salary  // Правильный вариант
}

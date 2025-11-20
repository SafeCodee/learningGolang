package main

import "fmt"

// TODO: Определи структуру Person с полями:
// - Name (string, экспортируемое)
// - Age (int, экспортируемое)
// - email (string, приватное)

// TODO: Определи структуру Address с полями:
// - City (string)
// - Country (string)

// TODO: Определи структуру Employee, которая:
// - Содержит поле Name (string)
// - Содержит поле Salary (float64)
// - ВСТРАИВАЕТ структуру Address (чтобы поля City и Country были доступны напрямую)

// TODO: Создай функцию-конструктор NewPerson, которая:
// - Принимает name (string), age (int), email (string)
// - Возвращает указатель на Person (*Person)
// - Инициализирует все поля переданными значениями

// TODO: Создай функцию-конструктор NewEmployee, которая:
// - Принимает name (string), salary (float64), city (string), country (string)
// - Возвращает указатель на Employee (*Employee)
// - Инициализирует все поля (включая встроенный Address)

// TODO: Создай функцию PrintPerson, которая:
// - Принимает указатель на Person
// - Выводит информацию в формате: "Name: <name>, Age: <age>, Email: <email>"

// TODO: Создай функцию PrintEmployee, которая:
// - Принимает указатель на Employee
// - Выводит информацию в формате: "Employee: <name>, Salary: <salary>, Location: <city>, <country>"

func main() {
	fmt.Println("=== Задание 1: Создание Person через литерал ===")
	// TODO: Создай переменную person1 типа Person с именованными полями:
	// Name = "Alice", Age = 30, email = "alice@example.com"

	// TODO: Выведи person1.Name и person1.Age

	fmt.Println("\n=== Задание 2: Создание Person через конструктор ===")
	// TODO: Создай person2 через NewPerson("Bob", 25, "bob@example.com")

	// TODO: Вызови функцию PrintPerson для person2

	fmt.Println("\n=== Задание 3: Zero value ===")
	// TODO: Создай переменную person3 типа Person БЕЗ инициализации (var person3 Person)

	// TODO: Выведи все поля person3 (Name, Age, email) - покажи zero values

	fmt.Println("\n=== Задание 4: Создание через указатель ===")
	// TODO: Создай person4 через &Person{} с полями Name = "Charlie", Age = 35

	// TODO: Измени Age у person4 на 36 через указатель

	// TODO: Выведи person4.Name и person4.Age

	fmt.Println("\n=== Задание 5: Встраивание структур ===")
	// TODO: Создай employee1 через NewEmployee("Diana", 75000.0, "Berlin", "Germany")

	// TODO: Вызови функцию PrintEmployee для employee1

	// TODO: Выведи employee1.City и employee1.Country напрямую (благодаря встраиванию)

	fmt.Println("\n=== Задание 6: Явный доступ к встроенной структуре ===")
	// TODO: Создай employee2 через литерал с именованными полями:
	// Name = "Eve", Salary = 80000.0
	// Address = Address{City: "Paris", Country: "France"}

	// TODO: Выведи employee2.Address.City через явное указание встроенной структуры

	fmt.Println("\n=== Задание 7: Анонимная структура ===")
	// TODO: Создай переменную temp с анонимной структурой, содержащей:
	// - поле Title (string) = "Temporary Data"
	// - поле Value (int) = 42

	// TODO: Выведи temp.Title и temp.Value

	fmt.Println("\n=== Задание 8: Сравнение структур ===")
	// TODO: Создай два Address с одинаковыми значениями:
	// addr1 := Address{City: "Moscow", Country: "Russia"}
	// addr2 := Address{City: "Moscow", Country: "Russia"}

	// TODO: Сравни addr1 == addr2 и выведи результат

	// TODO: Создай addr3 с другими значениями и сравни addr1 == addr3
}

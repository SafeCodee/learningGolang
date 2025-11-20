package main

import "fmt"

// TODO: Определи структуру Book с полями:
// - Title (string)
// - Author (string)
// - Pages (int)
// - CurrentPage (int)

// TODO: Создай метод Info() с value receiver, который выводит информацию о книге
// Формат: "Title by Author (Pages pages)"
// Пример: "Go Programming by John Doe (350 pages)"

// TODO: Создай метод Read(pages int) с pointer receiver
// Метод должен увеличивать CurrentPage на указанное количество страниц
// Если CurrentPage + pages превышает Pages, установи CurrentPage = Pages
// Выведи сообщение: "Read X pages. Current page: Y"

// TODO: Создай метод IsFinished() с value receiver, который возвращает bool
// Возвращает true если CurrentPage >= Pages

// TODO: Создай метод Progress() с value receiver, который возвращает float64
// Возвращает процент прочитанных страниц (CurrentPage / Pages * 100)

// TODO: Создай метод Reset() с pointer receiver
// Устанавливает CurrentPage в 0
// Выводит сообщение: "Book reset to the beginning"

// TODO: Определи структуру BankAccount с полями:
// - Owner (string)
// - Balance (float64)

// TODO: Создай метод Deposit(amount float64) с pointer receiver
// Добавляет amount к Balance
// Выводит: "Deposited $X. New balance: $Y"

// TODO: Создай метод Withdraw(amount float64) bool с pointer receiver
// Если amount <= Balance, вычитает amount из Balance, выводит успешное сообщение и возвращает true
// Если amount > Balance, выводит "Insufficient funds" и возвращает false

// TODO: Создай метод CheckBalance() с value receiver
// Выводит: "Owner's balance: $Balance"

// TODO: Определи новый тип Celsius на основе float64

// TODO: Создай метод ToFahrenheit() для типа Celsius с value receiver
// Формула: F = C * 9/5 + 32
// Возвращает Fahrenheit (можешь создать тип или использовать float64)

// TODO: Определи новый тип Fahrenheit на основе float64

// TODO: Создай метод ToCelsius() для типа Fahrenheit с value receiver
// Формула: C = (F - 32) * 5/9
// Возвращает Celsius

func main() {
	fmt.Println("=== Book Testing ===")
	// TODO: Создай книгу "The Go Programming Language" автора "Alan Donovan", 400 страниц

	// TODO: Вызови метод Info()

	// TODO: Прочитай 50 страниц

	// TODO: Выведи прогресс чтения (используй Printf с форматом %.2f для процентов)

	// TODO: Прочитай ещё 300 страниц

	// TODO: Проверь завершена ли книга (IsFinished), выведи результат

	// TODO: Прочитай ещё 100 страниц (должно ограничиться максимумом)

	// TODO: Проверь прогресс снова

	// TODO: Сбрось прогресс книги

	// TODO: Выведи текущий прогресс после сброса

	fmt.Println("\n=== Bank Account Testing ===")
	// TODO: Создай банковский счёт для "Alice" с балансом 1000.0

	// TODO: Проверь баланс

	// TODO: Внеси 500.0

	// TODO: Сними 200.0

	// TODO: Попытайся снять 2000.0 (должно не хватить средств)

	// TODO: Проверь финальный баланс

	fmt.Println("\n=== Temperature Conversion Testing ===")
	// TODO: Создай температуру 0 градусов Цельсия

	// TODO: Преобразуй в Фаренгейт и выведи результат
	// Формат: "0.00°C = 32.00°F"

	// TODO: Создай температуру 98.6 градусов Фаренгейта

	// TODO: Преобразуй в Цельсий и выведи результат
	// Формат: "98.60°F = 37.00°C"

	// TODO: Создай температуру 100°C, преобразуй в Фаренгейт
	// Выведи: "Boiling point: 100.00°C = 212.00°F"
}

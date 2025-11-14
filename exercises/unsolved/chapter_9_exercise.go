package main

import "fmt"

// TODO: Определи структуру Book с полями:
// - Title (string)
// - Author (string)
// - Pages (int)
// - CurrentPage (int)
type Book struct {
	Title       string
	Author      string
	Pages       int
	CurrentPage int
}

// TODO: Создай метод Info() с value receiver, который выводит информацию о книге
// Формат: "Title by Author (Pages pages)"
// Пример: "Go Programming by John Doe (350 pages)"
func (b Book) Info() {
	fmt.Printf("%s by %s (%d pages)\n", b.Title, b.Author, b.Pages)
}

// TODO: Создай метод Read(pages int) с pointer receiver
// Метод должен увеличивать CurrentPage на указанное количество страниц
// Если CurrentPage + pages превышает Pages, установи CurrentPage = Pages
// Выведи сообщение: "Read X pages. Current page: Y"
func (b *Book) Read(pages int) {
	if b.CurrentPage+pages >= b.Pages {
		b.CurrentPage = b.Pages
	} else {
		b.CurrentPage += pages
	}

	fmt.Printf("Read %d pages. Current page: %d\n", pages, b.CurrentPage)
}

// TODO: Создай метод IsFinished() с value receiver, который возвращает bool
// Возвращает true если CurrentPage >= Pages
func (b Book) isFinished() bool {
	return b.CurrentPage >= b.Pages
}

// TODO: Создай метод Progress() с value receiver, который возвращает float64
// Возвращает процент прочитанных страниц (CurrentPage / Pages * 100)
func (b Book) Progress() float64 {
	return float64(b.CurrentPage) / float64(b.Pages) * 100
}

// TODO: Создай метод Reset() с pointer receiver
// Устанавливает CurrentPage в 0
// Выводит сообщение: "Book reset to the beginning"
func (b *Book) Reset() {
	b.CurrentPage = 0
	fmt.Printf("Book reset to the beginning\n")
}

// TODO: Определи структуру BankAccount1 с полями:
// - Owner (string)
// - Balance (float64)
type BankAccount1 struct {
	Owner   string
	Balance float64
}

// TODO: Создай метод Deposit(amount float64) с pointer receiver
// Добавляет amount к Balance
// Выводит: "Deposited $X. New balance: $Y"
func (a *BankAccount1) Deposit(amount float64) {
	a.Balance += amount
	fmt.Printf("Deposited %v. New balance: %v\n", amount, a.Balance)
}

// TODO: Создай метод Withdraw(amount float64) bool с pointer receiver
// Если amount <= Balance, вычитает amount из Balance, выводит успешное сообщение и возвращает true
// Если amount > Balance, выводит "Insufficient funds" и возвращает false
func (a *BankAccount1) Withdraw(amount float64) bool {
	if amount <= a.Balance {
		a.Balance -= amount
		fmt.Println("Withdraw success")
		return true
	} else {
		fmt.Println("Insufficient funds")
		return false
	}
}

// TODO: Создай метод CheckBalance() с value receiver
// Выводит: "Owner's balance: $Balance"
func (a BankAccount1) CheckBalance() {
	fmt.Printf("Owner's balance: %v\n", a.Balance)
}

// TODO: Определи новый тип Celsius на основе float64
type Celsius float64

// TODO: Создай метод ToFahrenheit() для типа Celsius с value receiver
// Формула: F = C * 9/5 + 32
// Возвращает Fahrenheit (можешь создать тип или использовать float64)
func (c Celsius) ToFahrenheit() float64 {
	return float64(c)*9.0/5 + 32
}

// TODO: Определи новый тип Fahrenheit на основе float64
type Fahrenheit float64

// TODO: Создай метод ToCelsius() для типа Fahrenheit с value receiver
// Формула: C = (F - 32) * 5/9
// Возвращает Celsius
func (f Fahrenheit) ToCelsius() float64 {
	return (float64(f) - 32) * 5.0 / 9
}

func main() {
	fmt.Println("=== Book Testing ===")
	// TODO: Создай книгу "The Go Programming Language" автора "Alan Donovan", 400 страниц
	book := Book{
		Title:  "The Go Programming Language",
		Author: "Alan Donovan",
		Pages:  400,
	}
	// TODO: Вызови метод Info()
	book.Info()
	// TODO: Прочитай 50 страниц
	book.Read(50)
	// TODO: Выведи прогресс чтения (используй Printf с форматом %.2f для процентов)
	fmt.Printf("Progress %.2f\n", book.Progress())
	// TODO: Прочитай ещё 300 страниц
	book.Read(300)

	// TODO: Проверь завершена ли книга (IsFinished), выведи результат
	fmt.Printf("Book is finished %v\n", book.isFinished())

	// TODO: Прочитай ещё 100 страниц (должно ограничиться максимумом)
	book.Read(100)

	// TODO: Проверь прогресс снова
	fmt.Printf("Progress %.2f\n", book.Progress())

	// TODO: Сбрось прогресс книги
	book.Reset()
	// TODO: Выведи текущий прогресс после сброса
	fmt.Printf("Progress %.2f\n", book.Progress())

	fmt.Println("\n=== Bank Account Testing ===")
	// TODO: Создай банковский счёт для "Alice" с балансом 1000.0
	account := BankAccount1{
		Owner:   "Alice",
		Balance: 1000.0,
	}
	// TODO: Проверь баланс
	account.CheckBalance()
	// TODO: Внеси 500.0
	account.Deposit(500.0)
	// TODO: Сними 200.0
	account.Withdraw(200.0)
	// TODO: Попытайся снять 2000.0 (должно не хватить средств)
	account.Withdraw(2000.0)
	// TODO: Проверь финальный баланс
	account.CheckBalance()
	fmt.Println("\n=== Temperature Conversion Testing ===")
	// TODO: Создай температуру 0 градусов Цельсия
	celsius := Celsius(0.0)
	// TODO: Преобразуй в Фаренгейт и выведи результат
	// Формат: "0.00°C = 32.00°F"
	fmt.Printf("%.2f°C = %.2f°F\n", celsius, celsius.ToFahrenheit())
	// TODO: Создай температуру 98.6 градусов Фаренгейта
	fahrenheit := Fahrenheit(98.6)
	// TODO: Преобразуй в Цельсий и выведи результат
	// Формат: "98.60°F = 37.00°C"
	fmt.Printf("%.2f°F = %.2f°C\n", fahrenheit, fahrenheit.ToCelsius())
	// TODO: Создай температуру 100°C, преобразуй в Фаренгейт
	// Выведи: "Boiling point: 100.00°C = 212.00°F"
	boilingTemp := Celsius(100.0)
	fmt.Printf("Boiling point: %.2f°C = %.2f°F", boilingTemp, boilingTemp.ToFahrenheit())
}

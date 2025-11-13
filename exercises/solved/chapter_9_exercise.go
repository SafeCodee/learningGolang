package main

import "fmt"

// Структура Book с полями для отслеживания прогресса чтения
type Book struct {
	Title       string
	Author      string
	Pages       int
	CurrentPage int
}

// Info выводит информацию о книге (value receiver - только читает)
func (b Book) Info() {
	fmt.Printf("%s by %s (%d pages)\n", b.Title, b.Author, b.Pages)
}

// Read увеличивает текущую страницу (pointer receiver - изменяет структуру)
func (b *Book) Read(pages int) {
	b.CurrentPage += pages
	if b.CurrentPage > b.Pages {
		b.CurrentPage = b.Pages
	}
	fmt.Printf("Read %d pages. Current page: %d\n", pages, b.CurrentPage)
}

// IsFinished проверяет завершено ли чтение (value receiver - только читает)
func (b Book) IsFinished() bool {
	return b.CurrentPage >= b.Pages
}

// Progress возвращает процент прочитанных страниц (value receiver - вычисление)
func (b Book) Progress() float64 {
	return float64(b.CurrentPage) / float64(b.Pages) * 100
}

// Reset сбрасывает прогресс чтения (pointer receiver - изменяет структуру)
func (b *Book) Reset() {
	b.CurrentPage = 0
	fmt.Println("Book reset to the beginning")
}

// Структура BankAccount для управления балансом
type BankAccount struct {
	Owner   string
	Balance float64
}

// Deposit добавляет деньги на счёт (pointer receiver - изменяет баланс)
func (a *BankAccount) Deposit(amount float64) {
	a.Balance += amount
	fmt.Printf("Deposited $%.2f. New balance: $%.2f\n", amount, a.Balance)
}

// Withdraw снимает деньги со счёта (pointer receiver - изменяет баланс)
func (a *BankAccount) Withdraw(amount float64) bool {
	if amount <= a.Balance {
		a.Balance -= amount
		fmt.Printf("Withdrew $%.2f. New balance: $%.2f\n", amount, a.Balance)
		return true
	}
	fmt.Println("Insufficient funds")
	return false
}

// CheckBalance выводит текущий баланс (value receiver - только читает)
func (a BankAccount) CheckBalance() {
	fmt.Printf("%s's balance: $%.2f\n", a.Owner, a.Balance)
}

// Celsius - тип для температуры в градусах Цельсия
type Celsius float64

// ToFahrenheit конвертирует Цельсии в Фаренгейты
func (c Celsius) ToFahrenheit() Fahrenheit {
	return Fahrenheit(c*9.0/5 + 32)
}

// Fahrenheit - тип для температуры в градусах Фаренгейта
type Fahrenheit float64

// ToCelsius конвертирует Фаренгейты в Цельсии
func (f Fahrenheit) ToCelsius() Celsius {
	return Celsius((f - 32) * 5.0 / 9)
}

func main() {
	fmt.Println("=== Book Testing ===")
	book := Book{
		Title:  "The Go Programming Language",
		Author: "Alan Donovan",
		Pages:  400,
	}

	book.Info()
	book.Read(50)
	fmt.Printf("Progress: %.2f%%\n", book.Progress())

	book.Read(300)
	fmt.Printf("Book is finished: %v\n", book.IsFinished())

	book.Read(100)
	fmt.Printf("Progress: %.2f%%\n", book.Progress())

	book.Reset()
	fmt.Printf("Progress after reset: %.2f%%\n", book.Progress())

	fmt.Println("\n=== Bank Account Testing ===")
	account := BankAccount{
		Owner:   "Alice",
		Balance: 1000.0,
	}

	account.CheckBalance()
	account.Deposit(500.0)
	account.Withdraw(200.0)
	account.Withdraw(2000.0)
	account.CheckBalance()

	fmt.Println("\n=== Temperature Conversion Testing ===")
	celsius := Celsius(0.0)
	fmt.Printf("%.2f°C = %.2f°F\n", celsius, celsius.ToFahrenheit())

	fahrenheit := Fahrenheit(98.6)
	fmt.Printf("%.2f°F = %.2f°C\n", fahrenheit, fahrenheit.ToCelsius())

	boilingTemp := Celsius(100.0)
	fmt.Printf("Boiling point: %.2f°C = %.2f°F\n", boilingTemp, boilingTemp.ToFahrenheit())
}

package main

import (
	"fmt"
	"math"
	"strings"
)

// TODO: Импортируй необходимые пакеты

// Блок 1: Геометрические фигуры с интерфейсом Shape
// ===================================================

// TODO: Определи интерфейс Shape с двумя методами:
// - Area() float64
// - Perimeter() float64
type Shape interface {
	Area() float64
	Perimeter() float64
}

// TODO: Создай структуру Rectangle с полями Width и Height (float64)
type Rectangle struct {
	Width, Height float64
}

// TODO: Реализуй методы Area() и Perimeter() для Rectangle
// Area = Width * Height
// Perimeter = 2 * (Width + Height)

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// TODO: Создай структуру Circle с полем Radius (float64)
type Circle struct {
	Radius float64
}

// TODO: Реализуй методы Area() и Perimeter() для Circle
// Area = π * Radius²  (используй 3.14159)
// Perimeter = 2 * π * Radius
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// TODO: Создай структуру Triangle с полями A, B, C (float64) - длины сторон
type Triangle struct {
	A, B, C float64
}

// TODO: Реализуй методы Area() и Perimeter() для Triangle
// Perimeter = A + B + C
// Area: используй формулу Герона: √(s * (s-a) * (s-b) * (s-c)), где s = (a+b+c)/2
// Импортируй math для math.Sqrt()
func (t Triangle) Area() float64 {
	s := t.Perimeter() / 2
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}

func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

// TODO: Напиши функцию PrintShapeInfo(s Shape)
// которая принимает любую фигуру и выводит её площадь и периметр в формате:
// Area: %.2f, Perimeter: %.2f
func PrintShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

// Блок 2: Животные с интерфейсом Animal
// ======================================

// TODO: Определи интерфейс Animal с методами:
// - Speak() string
// - Type() string
type Animal interface {
	Speak() string
	Type() string
}

// TODO: Создай структуру Dog с полем Name (string)
type Dog struct {
	Name string
}

// TODO: Реализуй методы для Dog:
// - Speak() возвращает "Woof! Woof!"
// - Type() возвращает "Dog"

func (d Dog) Speak() string {
	return "Woof! Woof!"
}

func (d Dog) Type() string {
	return "Dog"
}

// TODO: Создай структуру Cat с полем Name (string)
type Cat struct {
	Name string
}

// TODO: Реализуй методы для Cat:
// - Speak() возвращает "Meow!"
// - Type() возвращает "Cat"
func (d Cat) Speak() string {
	return "Meow!"
}

func (d Cat) Type() string {
	return "Cat"
}

// TODO: Создай структуру Cow с полем Name (string)
type Cow struct {
	Name string
}

// TODO: Реализуй методы для Cow:
// - Speak() возвращает "Moo!"
// - Type() возвращает "Cow"
func (d Cow) Speak() string {
	return "Moo!"
}

func (d Cow) Type() string {
	return "Cow"
}

// TODO: Напиши функцию DescribeAnimal(a Animal)
// которая выводит: "[Type] says: [Speak()]"
// Например: "Dog says: Woof! Woof!"
func DescribeAnimal(a Animal) {
	fmt.Printf("%s says: %s\n", a.Type(), a.Speak())
}

// Блок 3: Реализация fmt.Stringer
// ================================

// TODO: Создай структуру Product с полями:
// - Name (string)
// - Price (float64)
// - Quantity (int)
type Product struct {
	Name     string
	Price    float64
	Quantity int
}

// TODO: Реализуй метод String() string для Product (реализует fmt.Stringer)
// Формат вывода: "Product: [Name], Price: $[Price], Qty: [Quantity]"
// Используй fmt.Sprintf для форматирования Price с 2 знаками после запятой
func (p Product) String() string {
	return fmt.Sprintf("Product: %s, Price: %.2f, Qty: %d", p.Name, p.Price, p.Quantity)
}

// Блок 4: Type Assertion и Type Switch
// =====================================

// TODO: Напиши функцию ProcessValue(v interface{})
// которая принимает любой тип и обрабатывает его с помощью type switch:
// - int: выводит "Integer: [значение]"
// - string: выводит "String: [значение]" (в верхнем регистре, используй strings.ToUpper)
// - bool: выводит "Boolean: [значение]"
// - []int: выводит "Int slice with [длина] elements"
// - Shape: вызывает PrintShapeInfo для этой фигуры
// - default: выводит "Unknown type: [тип]" (используй %T)

func ProcessValue(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
	case string:
		fmt.Printf("String: %s\n", strings.ToUpper(v.(string)))
	case bool:
		fmt.Printf("Boolean %t\n", v.(bool))
	case []int:
		fmt.Printf("Int slice with %d elements\n", len(v.([]int)))
	case Shape:
		PrintShapeInfo(v.(Shape))
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}

}

// Блок 5: Интерфейс с pointer receiver
// ====================================

// TODO: Создай интерфейс BankAccount1 с методами:
// - Deposit(amount float64)
// - Withdraw(amount float64) bool
// - Balance() float64
type BankAccount interface {
	Deposit(amount float64)
	Withdraw(amount float64) bool
	Balance() float64
}

// TODO: Создай структуру Account с полями:
// - Owner (string)
// - balance (float64) - приватное поле!
type Account struct {
	Owner   string
	balance float64
}

// TODO: Реализуй методы для Account (используй POINTER receiver!):
//   - Deposit(amount float64) - увеличивает balance на amount
//   - Withdraw(amount float64) bool - уменьшает balance на amount, если достаточно средств
//     Возвращает true если успешно, false если недостаточно средств
//   - Balance() float64 - возвращает текущий balance
func (a *Account) Deposit(amount float64) {
	a.balance += amount
}

func (a *Account) Withdraw(amount float64) bool {
	if amount <= a.balance {
		a.balance -= amount
		fmt.Println("Withdraw success")
		return true
	} else {
		fmt.Println("Insufficient funds")
		return false
	}
}

func (a *Account) Balance() float64 {
	return a.balance
}

// TODO: Напиши функцию TransferMoney(from, to BankAccount1, amount float64) bool
// которая снимает amount с from и кладёт на to
// Возвращает true если перевод успешен, false если недостаточно средств
func TransferMoney(from, to BankAccount, amount float64) bool {
	withdraw := from.Withdraw(amount)
	if withdraw {
		to.Deposit(amount)
	}

	return withdraw
}

func main() {
	// TODO: Тест Блок 1 - Геометрические фигуры
	// Создай Rectangle(5, 3), Circle(4), Triangle(3, 4, 5)
	// Вызови PrintShapeInfo для каждой фигуры
	rectangle := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 4}
	triangle := Triangle{
		A: 3,
		B: 4,
		C: 5,
	}
	PrintShapeInfo(rectangle)
	PrintShapeInfo(circle)
	PrintShapeInfo(triangle)

	// TODO: Тест Блок 2 - Животные
	// Создай Dog, Cat, Cow с разными именами
	// Вызови DescribeAnimal для каждого
	dog := Dog{Name: "Dog"}
	cat := Cat{Name: "Cat"}
	cow := Cow{Name: "Cow"}
	DescribeAnimal(dog)
	DescribeAnimal(cat)
	DescribeAnimal(cow)

	// TODO: Тест Блок 3 - Product со String()
	// Создай несколько Product
	// Выведи их через fmt.Println (должен вызваться String())
	product1 := Product{
		Name:     "Product 1",
		Price:    1.0,
		Quantity: 1,
	}
	product2 := Product{
		Name:     "Product 2",
		Price:    2.0,
		Quantity: 2,
	}

	fmt.Println(product1)
	fmt.Println(product2)

	// TODO: Тест Блок 4 - ProcessValue
	// Вызови ProcessValue с разными типами:
	// - int (42)
	// - string ("hello")
	// - bool (true)
	// - []int{1, 2, 3, 4, 5}
	// - одной из фигур (например Rectangle)
	// - float64 (3.14)
	ProcessValue(42)
	ProcessValue("hello")
	ProcessValue(true)
	ProcessValue([]int{1, 2, 3, 4, 5})
	ProcessValue(rectangle)
	ProcessValue(3.14)

	// TODO: Тест Блок 5 - BankAccount1
	// Создай два Account (используй &Account{...} чтобы получить указатель!)
	// Положи деньги на оба счета через Deposit
	// Попробуй перевести деньги между счетами через TransferMoney
	// Попробуй перевести больше чем есть на счёте (должно вернуть false)
	accountFrom := &Account{
		Owner:   "Owner1",
		balance: 0,
	}
	accountTo := &Account{
		Owner:   "Owner1",
		balance: 0,
	}

	accountFrom.Deposit(100)
	accountTo.Deposit(10)

	isTransfer := TransferMoney(accountFrom, accountTo, 50)
	fmt.Printf("Is transfered %t, from: %f, to: %f\n", isTransfer, accountFrom.Balance(), accountTo.Balance())
}

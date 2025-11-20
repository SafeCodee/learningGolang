package main

import "fmt"

// TODO: Импортируй необходимые пакеты (math, strings)

// Блок 1: Геометрические фигуры с интерфейсом Shape
// ===================================================

// TODO: Определи интерфейс Shape с двумя методами:
// - Area() float64
// - Perimeter() float64

// TODO: Создай структуру Rectangle с полями Width и Height (float64)

// TODO: Реализуй методы Area() и Perimeter() для Rectangle
// Area = Width * Height
// Perimeter = 2 * (Width + Height)

// TODO: Создай структуру Circle с полем Radius (float64)

// TODO: Реализуй методы Area() и Perimeter() для Circle
// Area = π * Radius²  (используй math.Pi)
// Perimeter = 2 * π * Radius

// TODO: Создай структуру Triangle с полями A, B, C (float64) - длины сторон

// TODO: Реализуй методы Area() и Perimeter() для Triangle
// Perimeter = A + B + C
// Area: используй формулу Герона: √(s * (s-a) * (s-b) * (s-c)), где s = (a+b+c)/2
// Импортируй math для math.Sqrt()

// TODO: Напиши функцию PrintShapeInfo(s Shape)
// которая принимает любую фигуру и выводит её площадь и периметр в формате:
// "Area: %.2f, Perimeter: %.2f"

// Блок 2: Животные с интерфейсом Animal
// ======================================

// TODO: Определи интерфейс Animal с методами:
// - Speak() string
// - Type() string

// TODO: Создай структуру Dog с полем Name (string)

// TODO: Реализуй методы для Dog:
// - Speak() возвращает "Woof! Woof!"
// - Type() возвращает "Dog"

// TODO: Создай структуру Cat с полем Name (string)

// TODO: Реализуй методы для Cat:
// - Speak() возвращает "Meow!"
// - Type() возвращает "Cat"

// TODO: Создай структуру Cow с полем Name (string)

// TODO: Реализуй методы для Cow:
// - Speak() возвращает "Moo!"
// - Type() возвращает "Cow"

// TODO: Напиши функцию DescribeAnimal(a Animal)
// которая выводит: "[Type] says: [Speak()]"
// Например: "Dog says: Woof! Woof!"

// Блок 3: Реализация fmt.Stringer
// ================================

// TODO: Создай структуру Product с полями:
// - Name (string)
// - Price (float64)
// - Quantity (int)

// TODO: Реализуй метод String() string для Product (реализует fmt.Stringer)
// Формат вывода: "Product: [Name], Price: $[Price], Qty: [Quantity]"
// Используй fmt.Sprintf для форматирования Price с 2 знаками после запятой

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

// Блок 5: Интерфейс с pointer receiver
// ====================================

// TODO: Создай интерфейс BankAccount с методами:
// - Deposit(amount float64)
// - Withdraw(amount float64) bool
// - Balance() float64

// TODO: Создай структуру Account с полями:
// - Owner (string)
// - balance (float64) - приватное поле!

// TODO: Реализуй методы для Account (используй POINTER receiver!):
//   - Deposit(amount float64) - увеличивает balance на amount
//   - Withdraw(amount float64) bool - уменьшает balance на amount, если достаточно средств
//     Возвращает true если успешно, false если недостаточно средств
//   - Balance() float64 - возвращает текущий balance

// TODO: Напиши функцию TransferMoney(from, to BankAccount, amount float64) bool
// которая снимает amount с from и кладёт на to
// Возвращает true если перевод успешен, false если недостаточно средств

func main() {
	fmt.Println("=== Блок 1: Геометрические фигуры ===")
	// TODO: Создай Rectangle(5, 3), Circle(4), Triangle(3, 4, 5)
	// TODO: Вызови PrintShapeInfo для каждой фигуры

	fmt.Println("\n=== Блок 2: Животные ===")
	// TODO: Создай Dog, Cat, Cow с разными именами
	// TODO: Вызови DescribeAnimal для каждого

	fmt.Println("\n=== Блок 3: Product со String() ===")
	// TODO: Создай несколько Product
	// TODO: Выведи их через fmt.Println (должен вызваться String())

	fmt.Println("\n=== Блок 4: ProcessValue ===")
	// TODO: Вызови ProcessValue с разными типами:
	// - int (42)
	// - string ("hello")
	// - bool (true)
	// - []int{1, 2, 3, 4, 5}
	// - одной из фигур (например Rectangle)
	// - float64 (3.14)

	fmt.Println("\n=== Блок 5: BankAccount ===")
	// TODO: Создай два Account (используй &Account{...} чтобы получить указатель!)
	// TODO: Положи деньги на оба счета через Deposit
	// TODO: Попробуй перевести деньги между счетами через TransferMoney
	// TODO: Попробуй перевести больше чем есть на счёте (должно вернуть false)
	// TODO: Выведи балансы обоих счетов
}

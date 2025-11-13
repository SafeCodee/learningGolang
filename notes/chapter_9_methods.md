# Глава 9: Методы (Methods)

## Что такое методы в Go?

**В Java:** методы определяются внутри класса:
```java
public class Person {
    private String name;

    public void greet() {  // метод класса
        System.out.println("Hello, " + name);
    }
}
```

**В Go:** методы определяются *вне* структуры, но привязываются к типу через **receiver** (получатель):
```go
type Person struct {
    Name string
}

func (p Person) Greet() {  // метод с receiver
    fmt.Println("Hello,", p.Name)
}
```

**Ключевое отличие:** В Go метод — это обычная функция с **специальным аргументом** (receiver) между `func` и именем метода.

---

## Синтаксис методов

```go
func (receiverName ReceiverType) MethodName(parameters) returnType {
    // тело метода
}
```

**Компоненты:**
- `receiverName` — имя переменной receiver (обычно короткое: `p`, `r`, `c`)
- `ReceiverType` — тип, к которому привязывается метод
- `MethodName` — имя метода (с заглавной = public, со строчной = private)

**Пример:**
```go
type Rectangle struct {
    Width, Height float64
}

// Метод для вычисления площади
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Использование
rect := Rectangle{Width: 10, Height: 5}
fmt.Println(rect.Area())  // 50
```

---

## Value Receiver vs Pointer Receiver

Это **критическая концепция** в Go! В отличие от Java, где объекты всегда передаются по ссылке, в Go нужно явно выбирать.

### Value Receiver (получатель по значению)

```go
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
```

**Что происходит:**
- Метод получает **копию** структуры
- Изменения внутри метода НЕ влияют на оригинал
- Используется для методов, которые **только читают** данные

**Аналог в Java:** как будто все поля `final` — можно читать, но не изменять оригинал.

**Пример:**
```go
type Counter struct {
    Count int
}

// Value receiver — НЕ изменит оригинал!
func (c Counter) Increment() {
    c.Count++  // изменяет КОПИЮ
}

func main() {
    counter := Counter{Count: 0}
    counter.Increment()
    fmt.Println(counter.Count)  // 0 — не изменилось!
}
```

### Pointer Receiver (получатель-указатель)

```go
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}
```

**Что происходит:**
- Метод получает **указатель** на структуру
- Изменения внутри метода **влияют на оригинал**
- Используется для методов, которые **изменяют** данные

**Аналог в Java:** обычное поведение методов в Java — изменения сохраняются.

**Пример:**
```go
type Counter struct {
    Count int
}

// Pointer receiver — изменит оригинал!
func (c *Counter) Increment() {
    c.Count++  // изменяет оригинал
}

func main() {
    counter := Counter{Count: 0}
    counter.Increment()
    fmt.Println(counter.Count)  // 1 — изменилось!
}
```

**Важно:** Go автоматически преобразует `counter.Increment()` в `(&counter).Increment()` — синтаксический сахар для удобства.

---

## Когда использовать каждый тип receiver?

### Используй **Pointer Receiver** если:

1. **Метод изменяет структуру:**
```go
func (c *Counter) Increment() {
    c.Count++
}
```

2. **Структура большая** (копирование дорого):
```go
type LargeStruct struct {
    Data [1000000]int
}

// Pointer receiver — не копируем миллион чисел
func (ls *LargeStruct) Process() {
    // ...
}
```

3. **Консистентность:** если хоть один метод использует pointer receiver, обычно **все методы** типа используют pointer receiver.

### Используй **Value Receiver** если:

1. **Метод только читает данные:**
```go
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
```

2. **Структура маленькая** (несколько примитивов):
```go
type Point struct {
    X, Y int
}

func (p Point) Distance() float64 {
    return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}
```

3. **Иммутабельность:** хочешь гарантировать что метод не изменит структуру.

**Java аналогия:**
- Pointer receiver ≈ обычный метод в Java
- Value receiver ≈ метод в immutable классе (String, Integer)

---

## Примеры с комментариями

### Пример 1: Банковский счёт

```go
package main

import "fmt"

type BankAccount struct {
    Owner   string
    Balance float64
}

// Value receiver — только читает
func (b BankAccount) DisplayBalance() {
    fmt.Printf("%s's balance: $%.2f\n", b.Owner, b.Balance)
}

// Pointer receiver — изменяет баланс
func (b *BankAccount) Deposit(amount float64) {
    b.Balance += amount
}

// Pointer receiver — изменяет баланс
func (b *BankAccount) Withdraw(amount float64) bool {
    if b.Balance >= amount {
        b.Balance -= amount
        return true
    }
    return false
}

func main() {
    account := BankAccount{Owner: "Alice", Balance: 100.0}

    account.DisplayBalance()  // Alice's balance: $100.00

    account.Deposit(50.0)
    account.DisplayBalance()  // Alice's balance: $150.00

    if account.Withdraw(30.0) {
        fmt.Println("Withdrawal successful")
    }
    account.DisplayBalance()  // Alice's balance: $120.00
}
```

### Пример 2: Геометрические фигуры

```go
package main

import (
    "fmt"
    "math"
)

type Circle struct {
    Radius float64
}

// Value receiver — вычисление (не изменяет структуру)
func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

// Value receiver — вычисление
func (c Circle) Circumference() float64 {
    return 2 * math.Pi * c.Radius
}

// Pointer receiver — изменяет радиус
func (c *Circle) Scale(factor float64) {
    c.Radius *= factor
}

func main() {
    circle := Circle{Radius: 5.0}

    fmt.Printf("Area: %.2f\n", circle.Area())              // 78.54
    fmt.Printf("Circumference: %.2f\n", circle.Circumference()) // 31.42

    circle.Scale(2.0)  // увеличиваем радиус в 2 раза
    fmt.Printf("New area: %.2f\n", circle.Area())          // 314.16
}
```

---

## Методы на встроенных типах

В Go можно создавать методы **только на типах, определённых в текущем пакете**. Нельзя добавить метод к `int`, но можно создать свой тип:

```go
package main

import "fmt"

// Создаём новый тип на основе int
type Temperature int

// Теперь можем добавить методы
func (t Temperature) Celsius() float64 {
    return float64(t)
}

func (t Temperature) Fahrenheit() float64 {
    return float64(t)*9.0/5.0 + 32
}

func main() {
    temp := Temperature(25)
    fmt.Printf("%.1f°C = %.1f°F\n", temp.Celsius(), temp.Fahrenheit())
    // 25.0°C = 77.0°F
}
```

**В Java:** нельзя расширить `Integer`, нужен wrapper класс. В Go — создаём type alias и добавляем методы.

---

## Методы vs Функции

**Когда использовать метод:**
- Операция логически связана с типом данных
- Хочешь использовать синтаксис `object.Method()`
- Готовишься к реализации интерфейса (об этом в главе 10)

**Когда использовать функцию:**
- Операция работает с несколькими типами
- Утилитарная функция без привязки к данным

**Пример:**
```go
// Метод — операция над конкретным типом
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Функция — утилита, работает с любыми числами
func Max(a, b float64) float64 {
    if a > b {
        return a
    }
    return b
}
```

---

## Важные моменты

1. **Имя receiver:** обычно короткое (1-2 буквы), соответствующее типу:
   - `Person` → `p`
   - `Rectangle` → `r` или `rect`
   - `BankAccount` → `b` или `ba`

2. **Консистентность:** все методы типа должны использовать одинаковое имя receiver.

3. **Автоматическая разыменовка:** Go автоматически преобразует:
   ```go
   p := &Person{Name: "Alice"}
   p.Greet()  // Go автоматически: (*p).Greet()

   p2 := Person{Name: "Bob"}
   p2.SomePointerMethod()  // Go автоматически: (&p2).SomePointerMethod()
   ```

4. **Nil receiver:** методы на pointer receiver можно вызвать даже если указатель nil (но нужна осторожность!):
   ```go
   var p *Person
   p.Greet()  // паника если внутри обращаешься к p.Name
   ```

---

## Сравнение с Java

| Аспект | Java | Go |
|--------|------|-----|
| **Определение** | Внутри класса | Вне структуры, с receiver |
| **Передача this** | Неявно | Явно через receiver |
| **Изменение объекта** | Всегда по ссылке | Выбор: value/pointer receiver |
| **Методы на примитивах** | Невозможно | Через type alias |
| **Перегрузка методов** | Есть | Нет (разные имена) |

**Java:**
```java
public class Rectangle {
    private double width, height;

    public double area() {  // метод внутри класса
        return width * height;
    }

    public void scale(double factor) {  // всегда изменяет объект
        this.width *= factor;
        this.height *= factor;
    }
}
```

**Go:**
```go
type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {  // value receiver — не изменяет
    return r.Width * r.Height
}

func (r *Rectangle) Scale(factor float64) {  // pointer receiver — изменяет
    r.Width *= factor
    r.Height *= factor
}
```

---

## Резюме

✅ **Методы в Go** — функции с receiver параметром
✅ **Value receiver** — для чтения (копия структуры)
✅ **Pointer receiver** — для изменения (указатель на оригинал)
✅ **Правило:** если изменяешь → pointer, если читаешь → value
✅ **Консистентность:** обычно все методы типа используют один вид receiver
✅ **Go автоматически** преобразует `value.PointerMethod()` и `pointer.ValueMethod()`

**Следующая глава:** Интерфейсы — как методы позволяют реализовывать полиморфизм в Go.

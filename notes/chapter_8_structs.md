# Глава 8: Структуры (Structs)

## Что такое структуры?

**Структура (struct)** в Go — это пользовательский тип данных, который группирует связанные поля (fields). Это аналог классов в Java, но **БЕЗ методов внутри** (методы определяются отдельно, это будет в главе 9).

### Сравнение с Java

```java
// Java
public class Person {
    private String name;
    private int age;

    // Конструктор, геттеры, сеттеры...
}
```

```go
// Go
type Person struct {
    Name string
    Age  int
}
// Методы определяются ОТДЕЛЬНО (глава 9)
```

**Ключевые отличия:**
- ❌ **Нет классов** — есть структуры
- ❌ **Нет конструкторов** — инициализация через литералы или функции-конструкторы
- ❌ **Нет геттеров/сеттеров** — доступ к полям напрямую (если экспортированы)
- ❌ **Нет наследования** — используется композиция через встраивание
- ✅ **Поля с большой буквы** экспортируются (public), с маленькой — приватные

---

## 1. Определение структуры

```go
// Базовое определение
type Person struct {
    Name string  // Экспортируемое поле (public)
    Age  int     // Экспортируемое поле
    city string  // Приватное поле (видно только в пакете)
}

// Можно группировать поля одного типа
type Rectangle struct {
    Width, Height float64
}

// Пустая структура (занимает 0 байт!)
type EmptyStruct struct{}
```

---

## 2. Создание экземпляров структуры

### 2.1. Литерал с именованными полями (рекомендуется)

```go
p1 := Person{
    Name: "Alice",
    Age:  30,
}
// Поле city не указано → получит zero value: ""
```

### 2.2. Литерал с позиционными значениями (НЕ рекомендуется)

```go
p2 := Person{"Bob", 25, "Moscow"}
// ДОЛЖНЫ указать ВСЕ поля в порядке определения
// Хрупкий код: добавление поля сломает код
```

### 2.3. Zero value (нулевое значение)

```go
var p3 Person
// p3.Name = "" (zero value для string)
// p3.Age = 0   (zero value для int)
// p3.city = "" (zero value для string)
```

### 2.4. Создание через указатель (`new`)

```go
p4 := new(Person)
// p4 — это *Person (указатель)
// Поля инициализированы zero values
fmt.Println(p4.Name) // "" — Go автоматически разыменовывает указатель
```

### 2.5. Создание через указатель с литералом

```go
p5 := &Person{
    Name: "Charlie",
    Age:  35,
}
// p5 — это *Person (указатель на Person)
```

**Когда использовать указатель?**
- Когда нужно модифицировать структуру внутри функции
- Когда структура большая (избегаем копирования)
- Когда хотим передать `nil` как "отсутствие значения"

---

## 3. Доступ к полям

```go
p := Person{Name: "Alice", Age: 30}

// Чтение
fmt.Println(p.Name) // Alice
fmt.Println(p.Age)  // 30

// Запись
p.Age = 31
p.Name = "Alicia"

// Через указатель
ptr := &p
ptr.Age = 32         // Go автоматически разыменовывает (*ptr).Age
fmt.Println(ptr.Name) // Alicia
```

---

## 4. Анонимные структуры

Структуры без имени типа — полезны для одноразовых данных.

```go
person := struct {
    Name string
    Age  int
}{
    Name: "John",
    Age:  40,
}

fmt.Println(person.Name) // John
```

**Когда использовать:**
- Короткие временные структуры
- Тестовые данные
- JSON маршалинг без создания типа

---

## 5. Встраивание структур (Embedding) — композиция вместо наследования

В Java есть наследование (`extends`), в Go — **композиция через встраивание**.

### Java наследование
```java
class Animal {
    String name;
}

class Dog extends Animal {
    String breed;
}

Dog dog = new Dog();
dog.name = "Rex"; // Унаследованное поле
```

### Go встраивание
```go
type Animal struct {
    Name string
}

type Dog struct {
    Animal      // Встроенная структура (без имени поля!)
    Breed string
}

dog := Dog{
    Animal: Animal{Name: "Rex"},
    Breed:  "Labrador",
}

// Доступ к полям встроенной структуры НАПРЯМУЮ
fmt.Println(dog.Name)  // Rex — "promoted field"
fmt.Println(dog.Breed) // Labrador

// Или через явное указание
fmt.Println(dog.Animal.Name) // Rex
```

**Важные моменты:**
- Поля встроенной структуры "поднимаются" (promoted) на уровень внешней структуры
- Это **не наследование** — это просто синтаксический сахар для композиции
- Можно встраивать несколько структур
- При конфликте имён полей нужно указывать явно: `dog.Animal.Name`

### Пример с несколькими встраиваниями

```go
type Address struct {
    City    string
    Country string
}

type Contact struct {
    Email string
    Phone string
}

type Employee struct {
    Name    string
    Address // Встраивание Address
    Contact // Встраивание Contact
}

emp := Employee{
    Name: "Alice",
    Address: Address{
        City:    "Moscow",
        Country: "Russia",
    },
    Contact: Contact{
        Email: "alice@example.com",
        Phone: "+7-123-456",
    },
}

// Поля "подняты" на уровень Employee
fmt.Println(emp.City)  // Moscow
fmt.Println(emp.Email) // alice@example.com
```

---

## 6. Теги структур (Struct Tags)

Метаданные для полей — используются библиотеками (например, `encoding/json`, `encoding/xml`).

```go
type User struct {
    ID        int    `json:"id"`
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Password  string `json:"-"` // Игнорировать при JSON маршалинге
}
```

**Синтаксис:** `` `key:"value"` ``

Подробнее о тегах будет в главе 19 (JSON).

---

## 7. Сравнение структур

Структуры можно сравнивать оператором `==`, если **все поля сравниваемые**.

```go
type Point struct {
    X, Y int
}

p1 := Point{X: 1, Y: 2}
p2 := Point{X: 1, Y: 2}
p3 := Point{X: 3, Y: 4}

fmt.Println(p1 == p2) // true
fmt.Println(p1 == p3) // false
```

**Несравниваемые типы:** `slice`, `map`, `function`

```go
type Container struct {
    Data []int
}

c1 := Container{Data: []int{1, 2}}
c2 := Container{Data: []int{1, 2}}
// c1 == c2 // ОШИБКА компиляции: slice не сравнивается
```

---

## 8. Функции-конструкторы (идиома Go)

Так как нет конструкторов, в Go принято создавать функции `NewXxx()`.

```go
type Person struct {
    Name string
    Age  int
}

// Конструктор — возвращает указатель на структуру
func NewPerson(name string, age int) *Person {
    return &Person{
        Name: name,
        Age:  age,
    }
}

// Использование
p := NewPerson("Alice", 30)
fmt.Println(p.Name) // Alice
```

**Зачем возвращать указатель?**
- Избегаем копирования большой структуры
- Позволяем возвращать `nil` при ошибке
- Идиоматично для Go

---

## 9. Экспорт полей

Правило видимости (как `public`/`private` в Java):
- **Заглавная буква** → экспортируется (public)
- **Строчная буква** → приватное (видно только в пакете)

```go
package mypackage

type Person struct {
    Name string // Экспортируемое
    age  int    // Приватное
}

// Другой пакет:
import "mypackage"

p := mypackage.Person{Name: "Alice"} // OK
p.age = 30 // ОШИБКА: age не экспортируется
```

---

## 10. Структуры vs Java классы — ключевые отличия

| Особенность | Java | Go |
|-------------|------|-----|
| Объявление | `class Person { }` | `type Person struct { }` |
| Конструктор | `public Person() { }` | Функция `NewPerson()` (идиома) |
| Методы | Внутри класса | Определяются отдельно (глава 9) |
| Наследование | `extends` | **НЕТ** — используется встраивание |
| Интерфейсы | `implements` явно | Неявная реализация (глава 10) |
| Геттеры/сеттеры | `getName()`, `setName()` | Прямой доступ к полям |
| Видимость | `public`, `private`, `protected` | Заглавная/строчная буква |
| Zero value | `null` по умолчанию | Инициализируются zero values |

---

## Итоги главы 8

✅ **Структуры** — это композитные типы данных (аналог классов без методов внутри)
✅ **Создание:** литералы, `new()`, `&Type{}`
✅ **Встраивание** (embedding) заменяет наследование — композиция вместо иерархии
✅ **Функции-конструкторы** `NewXxx()` — идиома Go для инициализации
✅ **Экспорт:** заглавная буква = public, строчная = private
✅ **Сравнение:** `==` работает, если все поля сравниваемые
✅ **Анонимные структуры** полезны для одноразовых данных

**Следующий шаг:** Глава 9 — Методы (как добавлять поведение к структурам)

---

## Примеры для практики

### Пример 1: Базовая структура
```go
package main

import "fmt"

type Book struct {
    Title  string
    Author string
    Pages  int
}

func main() {
    book := Book{
        Title:  "The Go Programming Language",
        Author: "Donovan & Kernighan",
        Pages:  380,
    }

    fmt.Printf("%s by %s (%d pages)\n", book.Title, book.Author, book.Pages)
}
```

### Пример 2: Встраивание структур
```go
package main

import "fmt"

type Address struct {
    City    string
    Country string
}

type Person struct {
    Name    string
    Address // Встраивание
}

func main() {
    p := Person{
        Name: "Alice",
        Address: Address{
            City:    "Berlin",
            Country: "Germany",
        },
    }

    // Доступ через promoted field
    fmt.Printf("%s lives in %s, %s\n", p.Name, p.City, p.Country)
}
```

### Пример 3: Функция-конструктор
```go
package main

import "fmt"

type Rectangle struct {
    Width  float64
    Height float64
}

func NewRectangle(width, height float64) *Rectangle {
    if width <= 0 || height <= 0 {
        return nil // Невалидные размеры
    }
    return &Rectangle{
        Width:  width,
        Height: height,
    }
}

func main() {
    r1 := NewRectangle(10, 5)
    if r1 != nil {
        fmt.Printf("Rectangle: %.1f x %.1f\n", r1.Width, r1.Height)
    }

    r2 := NewRectangle(-1, 5)
    if r2 == nil {
        fmt.Println("Invalid rectangle dimensions")
    }
}
```

# Go Quick Reference

Шпаргалка по Go для Java/Spring разработчика. Обновляется по мере прохождения глав.

---

## Основы (главы 1-3)

### Hello World
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

### Переменные
```go
// Объявление с типом
var name string = "Alice"

// Type inference
var age = 25

// Короткая форма (только внутри функций)
count := 10

// Множественное объявление
var x, y int = 1, 2
a, b := "hello", 42
```

### Константы
```go
const Pi = 3.14
const (
    StatusOK    = 200
    StatusError = 500
)
```

### Управляющие конструкции

**if:**
```go
if x > 0 {
    fmt.Println("positive")
} else if x < 0 {
    fmt.Println("negative")
} else {
    fmt.Println("zero")
}

// if с инициализацией
if err := doSomething(); err != nil {
    return err
}
```

**for (единственный цикл в Go):**
```go
// Классический for
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// Аналог while
for condition {
    // ...
}

// Бесконечный цикл
for {
    // ...
}

// range (для массивов, слайсов, map, строк)
for i, v := range slice {
    fmt.Printf("index=%d, value=%d\n", i, v)
}
```

**switch:**
```go
switch day {
case "Monday":
    fmt.Println("Start of week")
case "Friday":
    fmt.Println("TGIF")
default:
    fmt.Println("Regular day")
}

// switch без условия (замена if-else)
switch {
case x > 0:
    fmt.Println("positive")
case x < 0:
    fmt.Println("negative")
default:
    fmt.Println("zero")
}
```

---

## Функции (глава 4)

### Базовый синтаксис
```go
func add(x int, y int) int {
    return x + y
}

// Сокращённо (одинаковые типы)
func add(x, y int) int {
    return x + y
}
```

### Множественный возврат
```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

// Использование
result, err := divide(10, 2)
if err != nil {
    log.Fatal(err)
}
```

### Именованные результаты
```go
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return // "голый" return возвращает именованные переменные
}
```

### Вариадические параметры
```go
func sum(numbers ...int) int {
    total := 0
    for _, n := range numbers {
        total += n
    }
    return total
}

sum(1, 2, 3, 4) // 10
```

### Defer
```go
func readFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close() // выполнится в конце функции

    // работа с файлом
    return nil
}
```

### Функции высшего порядка
```go
// Функция как параметр
func apply(fn func(int) int, value int) int {
    return fn(value)
}

// Использование
result := apply(func(x int) int { return x * 2 }, 5) // 10
```

### Замыкания (Closures)
```go
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

c := counter()
fmt.Println(c()) // 1
fmt.Println(c()) // 2
```

---

## Массивы и слайсы (глава 5)

### Массивы (фиксированная длина)
```go
var arr [5]int              // [0 0 0 0 0]
numbers := [3]int{1, 2, 3}  // [1 2 3]
primes := [...]int{2, 3, 5} // автоматический размер
```

### Слайсы (динамическая длина)
```go
// Литерал
nums := []int{1, 2, 3, 4, 5}

// make
s := make([]int, 5)      // length=5, capacity=5
s := make([]int, 3, 10)  // length=3, capacity=10

// nil slice
var s []int
```

### Операции со слайсами
```go
// Append
s = append(s, 1)
s = append(s, 2, 3, 4)
s = append(s, otherSlice...)

// Slicing [low:high]
sub := nums[1:4]  // элементы с индекса 1 до 3 (4 не включая)
sub := nums[:3]   // от начала до 3
sub := nums[3:]   // от 3 до конца
sub := nums[:]    // весь слайс

// Copy
dst := make([]int, len(src))
copy(dst, src)

// Length и Capacity
len(s)  // количество элементов
cap(s)  // ёмкость базового массива
```

### Итерация
```go
for i, v := range nums {
    fmt.Printf("index=%d, value=%d\n", i, v)
}

// Только значения
for _, v := range nums {
    fmt.Println(v)
}

// Только индексы
for i := range nums {
    fmt.Println(i)
}
```

### Сравнение и поиск (Go 1.21+)
```go
import "slices"

slices.Equal(a, b)       // true если содержимое одинаковое
slices.Contains(s, x)    // true если x в слайсе
slices.Index(s, x)       // индекс элемента или -1
```

---

## Мапы (глава 6)

### Объявление и инициализация
```go
// Объявление (nil-мапа, нельзя использовать!)
var m map[string]int

// Создание с make
m := make(map[string]int)
m := make(map[string]int, 100)  // с hint для ёмкости

// Литерал
ages := map[string]int{
    "Alice": 25,
    "Bob":   30,
}
```

### Основные операции
```go
// Добавление/обновление
m["key"] = value

// Чтение
value := m["key"]  // 0 если ключа нет (zero value)

// Проверка существования ключа (важно!)
value, ok := m["key"]
if ok {
    // ключ существует
}

// Идиома
if value, ok := m["key"]; ok {
    // работа с value
}

// Удаление
delete(m, "key")

// Длина
len(m)
```

### Итерация
```go
// Ключ и значение
for key, value := range m {
    fmt.Printf("%s: %d\n", key, value)
}

// Только ключи
for key := range m {
    fmt.Println(key)
}

// Только значения
for _, value := range m {
    fmt.Println(value)
}
```

⚠️ **Порядок итерации случайный!** Go намеренно рандомизирует порядок при каждом запуске.

### Nil-мапа vs пустая мапа
```go
var m1 map[string]int        // nil-мапа
m2 := make(map[string]int)   // пустая мапа

fmt.Println(m1 == nil)  // true
fmt.Println(m2 == nil)  // false

len(m1)     // 0 (OK)
x := m1["key"]  // OK, вернёт zero value
m1["key"] = 1   // ❌ PANIC!

m2["key"] = 1   // ✅ OK
```

### Вложенные мапы
```go
// map[string]map[string]int
scores := map[string]map[string]int{
    "Alice": {
        "Math": 90,
        "Physics": 85,
    },
}

// Динамическое создание
scores := make(map[string]map[string]int)
scores["Bob"] = make(map[string]int)  // ❗ Нужно создать вложенную мапу
scores["Bob"]["Math"] = 75
```

---

## Указатели (глава 7)

### Основы указателей
```go
// Объявление указателя
var p *int

// Взятие адреса (&)
x := 42
p = &x  // p указывает на x

// Разыменование (*)
fmt.Println(*p)  // 42 (значение по адресу)
*p = 100         // изменяем значение x через указатель
```

### Передача по значению vs по ссылке
```go
// По значению (копируется)
func increment(n int) {
    n++  // изменяет копию
}

// По ссылке (через указатель)
func increment(n *int) {
    *n++  // изменяет оригинал
}

// Использование
x := 10
increment(x)   // x не изменится
increment(&x)  // x изменится на 11
```

### Создание указателей
```go
// С помощью new()
p := new(int)  // выделяет память, инициализирует zero value
*p = 42

// Через &
x := 42
p := &x

// Возврат указателя на локальную переменную (OK в Go!)
func createPointer(value int) *int {
    return &value  // переменная "убежит" на heap
}
```

### Zero value и nil
```go
var p *int
fmt.Println(p == nil)  // true

// ⚠️ Разыменование nil вызывает panic!
// fmt.Println(*p)  // PANIC!

// Безопасная проверка
if p != nil {
    fmt.Println(*p)
}
```

### Указатели и функции
```go
// Swap через указатели
func swap(a, b *int) {
    *a, *b = *b, *a
}

// Или через множественный возврат (идиоматичнее)
func swap(a, b int) (int, int) {
    return b, a
}
```

### Важно знать
- ✅ По умолчанию всё передаётся **по значению** (копируется)
- ✅ Указатели нужны для **изменения значений** в функциях
- ✅ **Нет арифметики указателей** (безопаснее чем C/C++)
- ✅ Всегда проверяй на **`nil`** перед использованием
- ✅ Слайсы/мапы уже содержат ссылки внутри (обычно не нужны указатели)

---

## Сравнение с Java

| Go | Java | Заметки |
|---|---|---|
| **Переменные** |
| `var x int = 10` | `int x = 10;` | |
| `x := 10` | - | Короткая форма (только в функциях) |
| **Функции** |
| `func add(x, y int) int` | `int add(int x, int y)` | Тип после имени |
| `func f() (int, error)` | - | Множественный возврат |
| `defer f()` | `try-finally` | Отложенное выполнение |
| **Массивы/Слайсы** |
| `[5]int` | `int[5]` | Размер — часть типа в Go |
| `[]int` | `ArrayList<Integer>` | Динамический массив |
| `make([]int, 0, 10)` | `new ArrayList<>(10)` | Создание с capacity |
| `append(s, x)` | `list.add(x)` | append возвращает новый слайс! |
| `s[1:4]` | `list.subList(1, 4)` | Оба создают view |
| `len(s)` | `list.size()` | |
| `slices.Equal(a, b)` | `list.equals(other)` | Go 1.21+ |
| **Мапы** |
| `map[string]int` | `Map<String, Integer>` | Встроенный тип в Go |
| `make(map[K]V)` | `new HashMap<>()` | |
| `m[key] = value` | `m.put(key, value)` | |
| `value := m[key]` | `value = m.get(key)` | Zero value если ключа нет |
| `value, ok := m[key]` | `m.containsKey(key)` | Проверка существования |
| `delete(m, key)` | `m.remove(key)` | |
| `len(m)` | `m.size()` | |
| `for k, v := range m` | `for (Entry<K,V> e : m.entrySet())` | Порядок случайный! |
| **Указатели** |
| `&x` | - | Взять адрес переменной |
| `*p` | - | Разыменовать (получить значение) |
| `*int` | - | Тип "указатель на int" |
| `func f(n *int)` | - | Передача по ссылке (явно) |
| `func f(n int)` | `void f(int n)` | Передача примитива по значению |
| `func f(obj *Type)` | `void f(Object obj)` | Передача объекта: явно vs неявно |
| `nil` | `null` | Отсутствие значения для указателя |
| **Структуры и методы** |
| `type Person struct {...}` | `class Person {...}` | Структура vs класс |
| `p.Name` | `p.getName()` | Прямой доступ к полям |
| `func (p Person) Method()` | `public void method()` | Value receiver vs метод класса |
| `func (p *Person) Method()` | `public void method()` | Pointer receiver (изменяет) |
| Композиция через встраивание | Наследование | Замена наследования |
| **Управление потоком** |
| `for` | `for`, `while`, `do-while` | В Go только for |
| `switch` (без break) | `switch` (нужен break) | |
| `for range` | `for-each`, `Stream.forEach` | |

---

## Полезные пакеты

```go
import "fmt"      // форматированный ввод/вывод
import "errors"   // создание ошибок
import "slices"   // операции со слайсами (Go 1.21+)
import "strings"  // работа со строками
import "strconv"  // конвертация строк
import "sort"     // сортировка слайсов и базовых типов
```

---

## Структуры (глава 8)

### Определение структуры
```go
type Person struct {
    Name string  // Экспортируемое (public)
    Age  int     // Экспортируемое
    city string  // Приватное (только в пакете)
}

// Группировка полей одного типа
type Rectangle struct {
    Width, Height float64
}
```

### Создание экземпляров
```go
// 1. Литерал с именованными полями (рекомендуется)
p1 := Person{
    Name: "Alice",
    Age:  30,
}

// 2. Литерал с позиционными значениями (не рекомендуется)
p2 := Person{"Bob", 25, "Moscow"}

// 3. Zero value
var p3 Person  // все поля получают zero values

// 4. Создание через new()
p4 := new(Person)  // возвращает *Person

// 5. Создание указателя через литерал
p5 := &Person{Name: "Charlie", Age: 35}  // *Person
```

### Доступ к полям
```go
p := Person{Name: "Alice", Age: 30}

// Чтение
fmt.Println(p.Name)  // Alice

// Запись
p.Age = 31

// Через указатель (автоматическое разыменование)
ptr := &p
ptr.Age = 32  // Go автоматически делает (*ptr).Age
```

### Встраивание структур (композиция)
```go
type Address struct {
    City    string
    Country string
}

type Employee struct {
    Name    string
    Salary  float64
    Address  // Встроенная структура (без имени поля!)
}

emp := Employee{
    Name:   "Alice",
    Salary: 75000,
    Address: Address{
        City:    "Berlin",
        Country: "Germany",
    },
}

// Promoted fields (доступ напрямую)
fmt.Println(emp.City)  // Berlin

// Или явно через встроенную структуру
fmt.Println(emp.Address.City)  // Berlin
```

### Функции-конструкторы (идиома)
```go
// Принято создавать функции NewXxx()
func NewPerson(name string, age int) *Person {
    return &Person{
        Name: name,
        Age:  age,
    }
}

// Использование
p := NewPerson("Alice", 30)
```

### Анонимные структуры
```go
// Без имени типа
person := struct {
    Name string
    Age  int
}{
    Name: "John",
    Age:  40,
}
```

### Сравнение структур
```go
type Point struct {
    X, Y int
}

p1 := Point{X: 1, Y: 2}
p2 := Point{X: 1, Y: 2}

fmt.Println(p1 == p2)  // true

// ⚠️ Не работает если есть slice, map, function
```

### Теги структур
```go
type User struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Password string `json:"-"`  // игнорировать при JSON
}
```

---

## Методы (глава 9)

### Определение методов
```go
// Value receiver (получатель по значению)
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Pointer receiver (получатель-указатель)
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}
```

### Синтаксис
```go
func (receiverName ReceiverType) MethodName(params) returnType {
    // тело метода
}
```

**Компоненты:**
- `receiverName` — имя переменной receiver (обычно 1-2 буквы)
- `ReceiverType` — тип, к которому привязывается метод
- Заглавная буква = public, строчная = private

### Value receiver vs Pointer receiver

**Value receiver (копия структуры):**
```go
func (c Counter) Increment() {
    c.Count++  // изменяет КОПИЮ, не влияет на оригинал
}
```

**Pointer receiver (указатель на оригинал):**
```go
func (c *Counter) Increment() {
    c.Count++  // изменяет ОРИГИНАЛ
}
```

### Когда использовать каждый тип

**Pointer receiver:**
- ✅ Метод изменяет структуру
- ✅ Структура большая (копирование дорого)
- ✅ Консистентность: если один метод pointer — обычно все pointer

**Value receiver:**
- ✅ Метод только читает данные
- ✅ Структура маленькая (несколько примитивов)
- ✅ Иммутабельность: гарантия что метод не изменит структуру

### Автоматическая разыменовка
```go
counter := Counter{Count: 0}
counter.Increment()  // Go автоматически: (&counter).Increment()

ptr := &Counter{Count: 0}
ptr.Increment()      // Go автоматически: (*ptr).Increment() для value receiver
```

### Методы на custom типах
```go
type Celsius float64

func (c Celsius) ToFahrenheit() Fahrenheit {
    return Fahrenheit(c*9.0/5 + 32)
}

temp := Celsius(100.0)
fmt.Println(temp.ToFahrenheit())  // 212
```

### Примеры
```go
type BankAccount struct {
    Owner   string
    Balance float64
}

// Value receiver — только читает
func (a BankAccount) CheckBalance() {
    fmt.Printf("%s's balance: $%.2f\n", a.Owner, a.Balance)
}

// Pointer receiver — изменяет баланс
func (a *BankAccount) Deposit(amount float64) {
    a.Balance += amount
}

// Pointer receiver — изменяет и возвращает bool
func (a *BankAccount) Withdraw(amount float64) bool {
    if amount <= a.Balance {
        a.Balance -= amount
        return true
    }
    return false
}
```

---

## Интерфейсы (глава 10)

### Определение интерфейса
```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
```

### Неявная реализация (duck typing)
```go
type Rectangle struct {
    Width, Height float64
}

// Rectangle автоматически реализует Shape
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}
```

**Ключевое отличие от Java:**
- Нет `implements` — тип реализует интерфейс, если имеет все его методы
- "Если ходит как утка и крякает как утка, то это утка"

### Полиморфизм
```go
func PrintShapeInfo(s Shape) {
    fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

rect := Rectangle{Width: 5, Height: 3}
circ := Circle{Radius: 4}

PrintShapeInfo(rect)  // ✅ Работает
PrintShapeInfo(circ)  // ✅ Работает
```

### Пустой интерфейс (interface{} или any)
```go
// Любой тип реализует пустой интерфейс
func PrintAnything(v interface{}) {
    fmt.Println(v)
}

// Go 1.18+ - псевдоним any
func PrintAnything(v any) {
    fmt.Println(v)
}

PrintAnything(42)       // int
PrintAnything("hello")  // string
PrintAnything([]int{})  // slice
```

### Type Assertion (приведение типа)
```go
var i interface{} = "hello"

// Небезопасное (panic если тип не тот)
s := i.(string)

// Безопасное (с проверкой)
s, ok := i.(string)
if ok {
    fmt.Println("Это строка:", s)
}
```

### Type Switch
```go
func ProcessValue(v interface{}) {
    switch val := v.(type) {  // Присваивание конкретного типа
    case int:
        fmt.Printf("Integer: %d\n", val)
    case string:
        fmt.Printf("String: %s\n", val)
    case bool:
        fmt.Printf("Boolean: %t\n", val)
    case Shape:
        PrintShapeInfo(val)
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
}
```

### Value vs Pointer Receiver в интерфейсах

**Важное правило:**

| Метод | Интерфейс реализуют |
|-------|---------------------|
| `func (t Type) Method()` | И `Type`, и `*Type` |
| `func (t *Type) Method()` | **ТОЛЬКО** `*Type` |

```go
type Incrementer interface {
    Increment()
}

type Counter struct {
    Value int
}

// Pointer receiver
func (c *Counter) Increment() {
    c.Value++
}

c := Counter{Value: 0}
// var inc Incrementer = c   // ❌ Ошибка!
var inc Incrementer = &c     // ✅ Только *Counter реализует
```

### Встроенные интерфейсы

**fmt.Stringer** — кастомное строковое представление:
```go
type Product struct {
    Name  string
    Price float64
}

func (p Product) String() string {
    return fmt.Sprintf("Product: %s, Price: $%.2f", p.Name, p.Price)
}

// fmt.Println автоматически использует String()
fmt.Println(product)  // Product: Laptop, Price: $999.00
```

**error** — обработка ошибок:
```go
type MyError struct {
    Code    int
    Message string
}

func (e MyError) Error() string {
    return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}
```

**io.Reader/Writer** — работа с потоками:
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}
```

### Композиция интерфейсов
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

// ReadWriter объединяет оба интерфейса
type ReadWriter interface {
    Reader
    Writer
}
```

### Лучшие практики

**1. Маленькие интерфейсы (1-3 метода):**
```go
// ✅ Хорошо
type Reader interface {
    Read(p []byte) (n int, err error)
}

// ❌ Плохо (слишком много методов)
type HugeInterface interface {
    Method1()
    Method2()
    Method3()
    Method4()
    Method5()
}
```

**2. Принимай интерфейсы, возвращай структуры:**
```go
// ✅ Хорошо
func ProcessData(r io.Reader) *Result { ... }

// ❌ Плохо
func ProcessData(f *os.File) io.Reader { ... }
```

**3. Определяй интерфейсы у потребителя:**
```go
// В package consumer
type Storage interface {
    Save(data string) error
}

func ProcessAndSave(s Storage, data string) {
    s.Save(data)
}

// В package implementations
type FileStorage struct { ... }
func (f FileStorage) Save(data string) error { ... }
```

### Проверка реализации интерфейса (compile-time)
```go
// Убедиться что Dog реализует Speaker
var _ Speaker = Dog{}
var _ Speaker = (*Dog)(nil)  // Для pointer receiver
```

### Nil интерфейс
```go
var s Speaker  // nil

if s == nil {
    fmt.Println("Interface is nil")
}

// s.Speak()  // panic!

// ⚠️ Интерфейс с nil-указателем != nil интерфейс:
var c *Counter  // nil указатель
var inc Incrementer = c  // интерфейс НЕ nil!
if inc == nil {
    // НЕ выполнится!
}
```

---

**Обновлено:** 2025-11-14 (главы 1-10 завершены)

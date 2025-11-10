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

**Обновлено:** 2025-11-10 (главы 1-7 завершены)

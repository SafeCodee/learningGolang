# Глава 4: Функции

## Базовый синтаксис

```go
func functionName(параметр тип) возвращаемыйТип {
    // тело функции
    return значение
}
```

### Простая функция

```go
func greet(name string) string {
    return "Hello, " + name
}

// Использование
message := greet("Alice")
```

### Функция без возврата значения

```go
func printMessage(msg string) {
    fmt.Println(msg)
}
```

В Go нет `void` — просто не указывается тип возврата.

## Множественный возврат значений

**Это одна из ключевых особенностей Go** (в Java такого нет).

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("деление на ноль")
    }
    return a / b, nil
}

// Использование
result, err := divide(10, 2)
if err != nil {
    fmt.Println("Ошибка:", err)
} else {
    fmt.Println("Результат:", result)
}
```

**Идиома Go:** функция возвращает `(результат, error)`.

## Именованные результаты

```go
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return  // "Голый" return — возвращает x и y
}
```

Переменные `x` и `y` объявлены в сигнатуре и автоматически возвращаются.

## Несколько параметров одного типа

```go
// Вместо: func add(x int, y int, z int)
func add(x, y, z int) int {
    return x + y + z
}
```

## Вариативные функции (Variadic Functions)

Аналог `...` в Java (varargs).

```go
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// Использование
result := sum(1, 2, 3, 4, 5)
```

**Как работает:**
- `numbers` — это слайс `[]int`
- Можно передать любое количество аргументов

## Функции как first-class citizens

В Go **функции — это значения** (как в Java с lambda).

### Функция как переменная

```go
func add(a, b int) int {
    return a + b
}

func main() {
    var operation func(int, int) int
    operation = add

    result := operation(5, 3)  // 8
}
```

### Анонимные функции (closures)

```go
// Анонимная функция
multiply := func(a, b int) int {
    return a * b
}

result := multiply(3, 4)  // 12
```

### Функции высшего порядка (Higher-Order Functions)

Функция, которая принимает или возвращает другую функцию.

```go
func apply(numbers []int, fn func(int) int) []int {
    result := make([]int, len(numbers))
    for i, num := range numbers {
        result[i] = fn(num)
    }
    return result
}

// Использование
numbers := []int{1, 2, 3, 4, 5}
doubled := apply(numbers, func(x int) int {
    return x * 2
})
// doubled = [2, 4, 6, 8, 10]
```

**Аналог в Java:** Stream API с lambda (`map(x -> x * 2)`).

## Замыкания (Closures)

Функция, которая "захватывает" переменные из внешней области видимости.

```go
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

// Использование
increment := counter()
fmt.Println(increment())  // 1
fmt.Println(increment())  // 2
fmt.Println(increment())  // 3
```

**Как в Java:** lambda с захватом переменных (effectively final).

## defer - отложенное выполнение

Функция будет вызвана **в конце текущей функции** (аналог `finally` в Java).

```go
func readFile() {
    file, err := os.Open("data.txt")
    if err != nil {
        return
    }
    defer file.Close()  // Выполнится в конце функции

    // Работа с файлом
    // ...
}
```

**Важно:** `defer` выполняется в обратном порядке (LIFO).

```go
func example() {
    defer fmt.Println("1")
    defer fmt.Println("2")
    defer fmt.Println("3")
}
// Вывод: 3, 2, 1
```

## Рекурсия

```go
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}
```

Работает как в Java.

## Отличия от Java

| Аспект | Java | Go |
|--------|------|-----|
| **Множественный возврат** | Нет (нужен класс/tuple) | Встроено: `func() (int, error)` |
| **Нет возврата** | `void` | Просто не указывается |
| **Varargs** | `int... numbers` | `numbers ...int` |
| **Lambda** | `(x) -> x * 2` | `func(x int) int { return x * 2 }` |
| **try-finally** | `try { } finally { }` | `defer cleanup()` |
| **Именованные результаты** | Нет | Есть |

## Паттерны использования функций

### 1. Обработка ошибок (идиоматично)

```go
func doSomething() error {
    if err := step1(); err != nil {
        return err
    }
    if err := step2(); err != nil {
        return err
    }
    return nil
}
```

### 2. Функции-конструкторы

```go
func NewUser(name string, age int) *User {
    return &User{
        Name: name,
        Age:  age,
    }
}
```

### 3. Опциональные параметры через функции

```go
type Config struct {
    Host string
    Port int
}

type Option func(*Config)

func WithHost(host string) Option {
    return func(c *Config) {
        c.Host = host
    }
}

func NewConfig(opts ...Option) *Config {
    cfg := &Config{Host: "localhost", Port: 8080}
    for _, opt := range opts {
        opt(cfg)
    }
    return cfg
}

// Использование
config := NewConfig(WithHost("example.com"))
```

## Важные моменты

✅ **Множественный возврат** — стандарт для (результат, ошибка)
✅ **Функции как значения** — можно передавать и возвращать
✅ **defer для cleanup** — аналог try-finally
✅ **Замыкания** — захват переменных из внешней области
✅ **Нет перегрузки функций** (в отличие от Java)

## Почему нет перегрузки?

В Java:
```java
int add(int a, int b) { }
double add(double a, double b) { }
```

В Go нужны **разные имена**:
```go
func addInt(a, b int) int { }
func addFloat(a, b float64) float64 { }
```

Философия Go: явность важнее краткости.

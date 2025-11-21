# Глава 11: Обработка ошибок

## Главное отличие от Java

**Java:**
```java
try {
    readFile("data.txt");
} catch (IOException e) {
    System.err.println("Ошибка: " + e.getMessage());
}
```

**Go:**
```go
data, err := readFile("data.txt")
if err != nil {
    fmt.Println("Ошибка:", err)
}
```

**Философия Go:** Ошибки — это обычные значения, а не исключения. Нет `try-catch`, ошибки возвращаются как второй результат функции.

---

## 1. Интерфейс `error`

В Go ошибка — это любой тип, реализующий встроенный интерфейс:

```go
type error interface {
    Error() string
}
```

Любая структура с методом `Error() string` автоматически становится ошибкой (duck typing).

---

## 2. Создание простых ошибок

### Вариант 1: `errors.New()`

```go
import "errors"

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("деление на ноль")
    }
    return a / b, nil
}

result, err := divide(10, 0)
if err != nil {
    fmt.Println("Ошибка:", err) // Ошибка: деление на ноль
}
```

### Вариант 2: `fmt.Errorf()` — с форматированием

```go
import "fmt"

func withdraw(balance, amount float64) (float64, error) {
    if amount > balance {
        return balance, fmt.Errorf("недостаточно средств: запрошено %.2f, доступно %.2f", amount, balance)
    }
    return balance - amount, nil
}

newBalance, err := withdraw(100, 150)
if err != nil {
    fmt.Println(err) // недостаточно средств: запрошено 150.00, доступно 100.00
}
```

**Аналог в Java:** `String.format()` + `new Exception(message)`

---

## 3. Проверка ошибок — идиома Go

**Стандартный паттерн:**
```go
result, err := someFunction()
if err != nil {
    // Обработка ошибки
    return err // или log.Fatal(err) или другая логика
}
// Работа с result (здесь err гарантированно nil)
```

**Важно:**
- Всегда проверяй ошибку сразу после вызова функции
- Если `err != nil`, обычно не используй другие возвращённые значения (они могут быть некорректными)
- Распространённая практика: если функция не может продолжить работу — возвращай ошибку выше по стеку

---

## 4. Кастомные ошибки — структуры с методом `Error()`

Когда нужна дополнительная информация об ошибке:

```go
type ValidationError struct {
    Field string
    Value interface{}
    Msg   string
}

// Реализация интерфейса error
func (e *ValidationError) Error() string {
    return fmt.Sprintf("поле '%s' (значение: %v): %s", e.Field, e.Value, e.Msg)
}

func validateAge(age int) error {
    if age < 0 || age > 150 {
        return &ValidationError{
            Field: "age",
            Value: age,
            Msg:   "возраст должен быть от 0 до 150",
        }
    }
    return nil
}

err := validateAge(200)
if err != nil {
    fmt.Println(err) // поле 'age' (значение: 200): возраст должен быть от 0 до 150
}
```

**Аналог в Java:** Создание собственного класса, наследующего `Exception`

---

## 5. Type Assertion для ошибок

Можно проверить конкретный тип ошибки:

```go
err := validateAge(-5)
if err != nil {
    if valErr, ok := err.(*ValidationError); ok {
        fmt.Println("Ошибка валидации поля:", valErr.Field)
    } else {
        fmt.Println("Другая ошибка:", err)
    }
}
```

**Аналог в Java:** `catch (SpecificException e)`

---

## 6. Sentinel Errors (предопределённые ошибки)

Экспортируемые переменные ошибок для сравнения:

```go
var ErrNotFound = errors.New("элемент не найден")
var ErrInvalidInput = errors.New("некорректный ввод")

func findUser(id int) (string, error) {
    if id < 0 {
        return "", ErrInvalidInput
    }
    // Поиск в БД...
    return "", ErrNotFound
}

name, err := findUser(123)
if err == ErrNotFound {
    fmt.Println("Создать нового пользователя?")
} else if err == ErrInvalidInput {
    fmt.Println("Проверь ID")
}
```

**Аналог в Java:** Константы типа `public static final Exception`

---

## 7. Wrapping Errors (оборачивание ошибок) — Go 1.13+

### Проблема:
```go
func processFile(filename string) error {
    _, err := os.Open(filename)
    if err != nil {
        return errors.New("не удалось обработать файл") // Потеряли исходную ошибку!
    }
    return nil
}
```

### Решение: `fmt.Errorf()` с `%w`
```go
func processFile(filename string) error {
    _, err := os.Open(filename)
    if err != nil {
        return fmt.Errorf("не удалось обработать файл %s: %w", filename, err)
    }
    return nil
}

err := processFile("data.txt")
if err != nil {
    fmt.Println(err)
    // Вывод: не удалось обработать файл data.txt: open data.txt: no such file or directory
}
```

### Распаковка: `errors.Unwrap()`, `errors.Is()`, `errors.As()`

```go
import "errors"

err := processFile("data.txt")

// Проверка на конкретную ошибку (аналог instanceof)
if errors.Is(err, os.ErrNotExist) {
    fmt.Println("Файл не существует")
}

// Извлечение кастомной ошибки из цепочки
var valErr *ValidationError
if errors.As(err, &valErr) {
    fmt.Println("Ошибка валидации:", valErr.Field)
}
```

**Аналог в Java:** `Throwable.getCause()`, но в Go более явно и безопасно

---

## 8. `panic` и `recover` — экстремальные случаи

### `panic` — аварийная остановка программы

```go
func mustConnect(url string) {
    if url == "" {
        panic("URL не может быть пустым!")
    }
    // Подключение...
}
```

**Когда использовать:**
- Только для **критических** ошибок (например, невозможность инициализации при старте)
- Аналог `throw new RuntimeException()` в Java, но используется **гораздо реже**

### `recover` — перехват паники

```go
func safeExecute(fn func()) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Восстановление после паники:", r)
        }
    }()
    fn()
}

safeExecute(func() {
    panic("что-то пошло не так!")
})
// Вывод: Восстановление после паники: что-то пошло не так!
```

**Важно:**
- `recover()` работает **только** внутри `defer`
- Используй редко (только для graceful shutdown или границ модулей)
- **Не используй** `panic/recover` для обычной обработки ошибок!

**Аналог в Java:** `try-catch`, но в Go это anti-pattern для обычного flow

---

## 9. Множественные ошибки (Go 1.20+)

```go
import "errors"

func validateUser(name string, age int) error {
    var errs []error

    if name == "" {
        errs = append(errs, errors.New("имя не может быть пустым"))
    }
    if age < 18 {
        errs = append(errs, errors.New("возраст должен быть >= 18"))
    }

    if len(errs) > 0 {
        return errors.Join(errs...)
    }
    return nil
}

err := validateUser("", 15)
if err != nil {
    fmt.Println(err)
    // Вывод (многострочный):
    // имя не может быть пустым
    // возраст должен быть >= 18
}
```

---

## 10. Идиомы обработки ошибок в Go

### ✅ Хорошо:
```go
file, err := os.Open("data.txt")
if err != nil {
    return fmt.Errorf("открытие файла: %w", err)
}
defer file.Close()

data, err := io.ReadAll(file)
if err != nil {
    return fmt.Errorf("чтение файла: %w", err)
}
```

### ❌ Плохо (игнорирование ошибок):
```go
file, _ := os.Open("data.txt") // НЕ делай так!
```

### ❌ Плохо (злоупотребление panic):
```go
func getUser(id int) User {
    user, err := db.Query(id)
    if err != nil {
        panic(err) // Используй return err вместо panic!
    }
    return user
}
```

---

## Сравнение с Java

| Концепция | Java | Go |
|-----------|------|-----|
| Обработка ошибок | `try-catch-finally` | явный возврат `error`, `defer` |
| Исключения | Checked/Unchecked | Нет концепции исключений |
| Создание ошибки | `throw new Exception("msg")` | `errors.New("msg")` |
| Цепочка ошибок | `Throwable.getCause()` | `fmt.Errorf("context: %w", err)` |
| Критические ошибки | `RuntimeException` | `panic` (используется редко) |
| Восстановление | `catch (Exception e)` | `recover()` в `defer` |
| Кастомная ошибка | `class MyException extends Exception` | Структура с методом `Error()` |

---

## Ключевые выводы

1. **Ошибки — это значения, а не исключения**
2. **Всегда проверяй `err != nil` сразу после вызова**
3. **Возвращай ошибки явно**, не скрывай их
4. **Используй `fmt.Errorf()` с `%w`** для добавления контекста
5. **`panic/recover` — только для критических случаев**, не для обычного flow
6. **Создавай кастомные ошибки** (структуры) когда нужна доп. информация
7. **Sentinel errors** — для предопределённых ошибок, которые нужно сравнивать

**Философия:** Go заставляет явно думать об ошибках на каждом шаге, что делает код более надёжным и предсказуемым.

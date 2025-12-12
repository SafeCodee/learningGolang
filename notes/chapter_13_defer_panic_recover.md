# Глава 13: Defer, Panic, Recover

## Введение

В этой главе разберём три важных механизма управления потоком выполнения в Go:
- **defer** — отложенное выполнение (аналог `finally` в Java, но мощнее)
- **panic** — аварийное завершение (похоже на `throw` исключения в Java)
- **recover** — восстановление после паники (похоже на `catch` в Java)

Важно понимать: в Go **ошибки обрабатываются через возврат значений** (error), а panic/recover используются **только для исключительных ситуаций** (критические ошибки, от которых невозможно восстановиться).

---

## 1. Defer — отложенное выполнение

### Что это такое?

`defer` откладывает выполнение функции до момента завершения окружающей функции. Отложенные вызовы выполняются **всегда**, даже если функция завершилась с ошибкой или паникой.

### Базовый синтаксис

```go
func example() {
    defer fmt.Println("Выполнится в конце")
    fmt.Println("Выполнится сначала")
}
// Вывод:
// Выполнится сначала
// Выполнится в конце
```

### Сравнение с Java

**Java (try-finally):**
```java
public void readFile() {
    FileReader reader = new FileReader("file.txt");
    try {
        // работа с файлом
    } finally {
        reader.close(); // всегда выполнится
    }
}
```

**Go (defer):**
```go
func readFile() {
    file, err := os.Open("file.txt")
    if err != nil {
        return
    }
    defer file.Close() // всегда выполнится при выходе из функции

    // работа с файлом
}
```

**Преимущества defer:**
- Код закрытия ресурса находится **рядом с открытием** (лучше читаемость)
- Не нужно дублировать `close()` в каждой ветке return
- Работает даже при панике

---

## 2. Порядок выполнения defer

### LIFO (Last In, First Out) — стек

Отложенные вызовы выполняются в **обратном порядке** (как стек):

```go
func main() {
    defer fmt.Println("1")
    defer fmt.Println("2")
    defer fmt.Println("3")
    fmt.Println("Основной код")
}
// Вывод:
// Основной код
// 3
// 2
// 1
```

**Почему LIFO?** Чтобы ресурсы закрывались в обратном порядке их открытия (последний открытый закрывается первым).

### Пример с ресурсами

```go
func processData() {
    db := openDatabase()
    defer db.Close() // закроется последним

    file := openFile()
    defer file.Close() // закроется первым

    // работа с db и file
}
// Порядок закрытия:
// 1. file.Close()
// 2. db.Close()
```

---

## 3. Defer и аргументы функции

### Важно! Аргументы вычисляются сразу

Аргументы отложенной функции **вычисляются в момент вызова defer**, а не в момент выполнения:

```go
func main() {
    x := 10
    defer fmt.Println("defer:", x) // x = 10 (зафиксировано здесь)
    x = 20
    fmt.Println("main:", x) // x = 20
}
// Вывод:
// main: 20
// defer: 10
```

### Замыкание (closure) для доступа к актуальным значениям

Если нужно получить **финальное значение**, используй анонимную функцию:

```go
func main() {
    x := 10
    defer func() {
        fmt.Println("defer:", x) // x = 20 (актуальное значение)
    }()
    x = 20
    fmt.Println("main:", x)
}
// Вывод:
// main: 20
// defer: 20
```

---

## 4. Практические примеры defer

### Работа с файлами

```go
func writeToFile(filename string, data []byte) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close() // гарантированно закроется

    _, err = file.Write(data)
    return err
}
```

### Разблокировка мьютекса

```go
var mu sync.Mutex

func criticalSection() {
    mu.Lock()
    defer mu.Unlock() // гарантированно разблокируется

    // критическая секция
}
```

### Измерение времени выполнения

```go
func measureTime(name string) func() {
    start := time.Now()
    return func() {
        fmt.Printf("%s took %v\n", name, time.Since(start))
    }
}

func slowFunction() {
    defer measureTime("slowFunction")() // обрати внимание на ()()

    time.Sleep(2 * time.Second)
}
```

---

## 5. Panic — аварийное завершение

### Что это такое?

`panic` вызывает **немедленное аварийное завершение** программы. Похоже на `throw new RuntimeException()` в Java, но в Go используется **редко**.

### Когда использовать panic?

**В Go panic используется ТОЛЬКО для критических ошибок:**
- Программная ошибка (баг), от которой невозможно восстановиться
- Невозможность продолжить работу (нет памяти, критический ресурс недоступен)

**НЕ используй panic для обычных ошибок!** Используй возврат `error`.

### Базовый синтаксис

```go
func divide(a, b int) int {
    if b == 0 {
        panic("деление на ноль") // аварийное завершение
    }
    return a / b
}

func main() {
    result := divide(10, 0) // программа аварийно завершится
    fmt.Println(result)     // не выполнится
}
```

### Что происходит при panic?

1. Выполнение текущей функции **останавливается**
2. Все **отложенные defer** в текущей функции **выполняются**
3. Управление возвращается в вызывающую функцию, процесс повторяется
4. Если panic не перехвачен — программа завершается с выводом stack trace

---

## 6. Recover — восстановление после паники

### Что это такое?

`recover` позволяет **перехватить panic** и продолжить работу программы. Работает **только внутри defer**.

### Базовый синтаксис

```go
func safeDivide(a, b int) (result int) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Перехвачена паника:", r)
            result = 0 // можно установить значение по умолчанию
        }
    }()

    return a / b // может вызвать панику при b = 0
}

func main() {
    result := safeDivide(10, 0)
    fmt.Println("Результат:", result) // программа продолжит работу
}
// Вывод:
// Перехвачена паника: runtime error: integer divide by zero
// Результат: 0
```

### Как работает recover?

- `recover()` возвращает значение, переданное в `panic()` (или `nil`, если паники не было)
- Работает **только внутри defer**
- Останавливает распространение паники

---

## 7. Сравнение с Java try-catch-finally

### Java

```java
public int divide(int a, int b) {
    try {
        return a / b;
    } catch (ArithmeticException e) {
        System.out.println("Ошибка: " + e.getMessage());
        return 0;
    } finally {
        System.out.println("Всегда выполнится");
    }
}
```

### Go (эквивалент)

```go
func divide(a, b int) (result int) {
    defer func() {
        fmt.Println("Всегда выполнится") // finally

        if r := recover(); r != nil {    // catch
            fmt.Println("Ошибка:", r)
            result = 0
        }
    }()

    return a / b // может вызвать panic
}
```

### Ключевые отличия

| Java | Go |
|------|-----|
| `try` — блок кода | Обычная функция |
| `catch` — перехват исключения | `recover()` в defer |
| `finally` — всегда выполнится | `defer` без recover |
| Исключения для обычных ошибок | `error` для обычных ошибок |
| `throw` для любых ошибок | `panic` только для критических |

---

## 8. Идиоматичный Go: когда использовать что?

### Обычные ошибки → error

```go
func readConfig(filename string) (*Config, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, err // возвращаем ошибку
    }

    var config Config
    if err := json.Unmarshal(data, &config); err != nil {
        return nil, err // возвращаем ошибку
    }

    return &config, nil
}
```

### Критические ошибки → panic

```go
func mustLoadConfig(filename string) *Config {
    config, err := readConfig(filename)
    if err != nil {
        panic("не могу загрузить конфигурацию: " + err.Error())
        // программа не может работать без конфига
    }
    return config
}
```

### Восстановление в горутинах → recover

```go
func safeGoroutine(task func()) {
    go func() {
        defer func() {
            if r := recover(); r != nil {
                fmt.Println("Горутина упала:", r)
                // логируем, но не роняем всю программу
            }
        }()

        task() // может вызвать панику
    }()
}
```

---

## 9. Примеры из реальной практики

### Пример 1: Безопасная работа с индексами массива

```go
func getElement(slice []int, index int) (value int, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("индекс вне диапазона: %v", r)
        }
    }()

    value = slice[index] // может вызвать панику
    return value, nil
}

func main() {
    numbers := []int{1, 2, 3}
    value, err := getElement(numbers, 10)
    if err != nil {
        fmt.Println("Ошибка:", err)
    } else {
        fmt.Println("Значение:", value)
    }
}
```

### Пример 2: Логирование выхода из функции

```go
func complexOperation(id int) (err error) {
    defer func() {
        if err != nil {
            fmt.Printf("Операция %d завершилась с ошибкой: %v\n", id, err)
        } else {
            fmt.Printf("Операция %d успешно завершена\n", id)
        }
    }()

    // сложная логика
    return nil
}
```

### Пример 3: Откат транзакции БД

```go
func transferMoney(db *sql.DB, from, to int, amount float64) error {
    tx, err := db.Begin()
    if err != nil {
        return err
    }

    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
            panic(r) // пробрасываем панику дальше
        }
    }()

    // операции с БД
    if err := debit(tx, from, amount); err != nil {
        tx.Rollback()
        return err
    }

    if err := credit(tx, to, amount); err != nil {
        tx.Rollback()
        return err
    }

    return tx.Commit()
}
```

---

## 10. Частые ошибки

### ❌ Неправильно: recover без defer

```go
func wrong() {
    recover() // НЕ РАБОТАЕТ! recover только в defer
    panic("ошибка")
}
```

### ✅ Правильно: recover в defer

```go
func correct() {
    defer func() {
        recover() // работает
    }()
    panic("ошибка")
}
```

### ❌ Неправильно: defer в цикле для ресурсов

```go
func processFiles(files []string) {
    for _, filename := range files {
        f, _ := os.Open(filename)
        defer f.Close() // ВСЕ файлы закроются только в конце функции!
    }
}
```

### ✅ Правильно: отдельная функция или ручное закрытие

```go
func processFiles(files []string) {
    for _, filename := range files {
        processFile(filename) // defer внутри processFile
    }
}

func processFile(filename string) {
    f, _ := os.Open(filename)
    defer f.Close() // закроется сразу после обработки
    // обработка файла
}
```

---

## Резюме

| Механизм | Назначение | Аналог в Java |
|----------|-----------|---------------|
| **defer** | Отложенное выполнение (закрытие ресурсов, cleanup) | `finally` |
| **panic** | Критические ошибки, невозможно продолжить | `throw RuntimeException` |
| **recover** | Перехват паники, восстановление | `catch` |

**Главное правило:**
- Обычные ошибки → `error` (возврат значения)
- Критические ошибки → `panic` (редко!)
- Восстановление → `recover` в `defer`

**Паттерны использования:**
- `defer` для закрытия ресурсов (файлы, соединения, мьютексы)
- `defer` для измерения времени, логирования
- `panic` только для программных ошибок
- `recover` в горутинах, чтобы не ронять всю программу
- `recover` для защиты публичных API от паник в пользовательском коде

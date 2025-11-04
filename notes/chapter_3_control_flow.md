# Глава 3: Управляющие конструкции

## if - условный оператор

### Базовый синтаксис

```go
if x > 10 {
    fmt.Println("x больше 10")
}
```

**Важно:** скобки `()` вокруг условия **не нужны**!

### if-else

```go
if x > 10 {
    fmt.Println("Больше 10")
} else {
    fmt.Println("Меньше или равно 10")
}
```

### if-else if-else

```go
if score >= 90 {
    fmt.Println("A")
} else if score >= 80 {
    fmt.Println("B")
} else {
    fmt.Println("C")
}
```

### if с инициализацией (короткое объявление)

```go
// Переменная err доступна только внутри if
if err := doSomething(); err != nil {
    fmt.Println("Ошибка:", err)
}
```

**Это идиоматично в Go** и часто используется для обработки ошибок.

## for - единственный цикл в Go

В Go **нет while, do-while** — только `for` со всеми вариантами.

### 1. Классический for (как в Java)

```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

### 2. Как while

```go
i := 0
for i < 10 {
    fmt.Println(i)
    i++
}
```

### 3. Бесконечный цикл

```go
for {
    // Бесконечный цикл
    // Выход через break
}
```

### 4. range (итерация по коллекциям)

```go
numbers := []int{1, 2, 3, 4, 5}

// С индексом и значением
for index, value := range numbers {
    fmt.Println(index, value)
}

// Только значения (игнорируем индекс)
for _, value := range numbers {
    fmt.Println(value)
}

// Только индексы
for index := range numbers {
    fmt.Println(index)
}
```

## switch - выбор варианта

### Базовый switch

```go
day := 3

switch day {
case 1:
    fmt.Println("Понедельник")
case 2:
    fmt.Println("Вторник")
case 3:
    fmt.Println("Среда")
default:
    fmt.Println("Другой день")
}
```

**Важно:** в Go **не нужен break** — выход автоматический после каждого case!

### Множественные значения в case

```go
switch day {
case 1, 2, 3, 4, 5:
    fmt.Println("Будний день")
case 6, 7:
    fmt.Println("Выходной")
}
```

### switch с условиями (без выражения)

```go
score := 85

switch {
case score >= 90:
    fmt.Println("A")
case score >= 80:
    fmt.Println("B")
case score >= 70:
    fmt.Println("C")
default:
    fmt.Println("F")
}
```

Это аналог множественных `if-else if`.

### switch с инициализацией

```go
switch result := calculate(); {
case result > 100:
    fmt.Println("Большое значение")
case result > 50:
    fmt.Println("Среднее значение")
default:
    fmt.Println("Малое значение")
}
```

### fallthrough (принудительное проваливание)

```go
switch x {
case 1:
    fmt.Println("Один")
    fallthrough  // Выполнит следующий case
case 2:
    fmt.Println("Два")
}
```

**Используется редко**, так как по умолчанию нет проваливания.

## break и continue

### break - выход из цикла

```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break  // Выход из цикла
    }
    fmt.Println(i)
}
```

### continue - пропуск итерации

```go
for i := 0; i < 10; i++ {
    if i%2 == 0 {
        continue  // Пропускаем чётные
    }
    fmt.Println(i)  // Печатаем только нечётные
}
```

## Отличия от Java

| Аспект | Java | Go |
|--------|------|-----|
| **if скобки** | `if (x > 10) { }` | `if x > 10 { }` |
| **Циклы** | for, while, do-while | Только for (все варианты) |
| **switch break** | Нужен break в каждом case | Автоматически (без break) |
| **switch fallthrough** | По умолчанию | Нужно явно указать fallthrough |
| **range** | for-each: `for (int x : list)` | `for _, x := range list` |
| **Короткое объявление в if** | Нет | `if err := foo(); err != nil` |

## Важные моменты

✅ **Нет скобок** вокруг условия в `if`
✅ **Один цикл for** заменяет for, while, do-while
✅ **switch без break** — автоматический выход
✅ **range для итерации** — удобнее foreach из Java
✅ **`_` для игнорирования** значений (blank identifier)

## Идиоматичный Go код

**Проверка ошибок с коротким объявлением:**
```go
if err := file.Write(data); err != nil {
    return err
}
```

**Итерация только по значениям:**
```go
for _, value := range collection {
    process(value)
}
```

**switch без выражения вместо if-else if:**
```go
switch {
case x > 100:
    // ...
case x > 50:
    // ...
}
```

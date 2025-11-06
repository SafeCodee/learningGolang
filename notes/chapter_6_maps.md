# Глава 6: Мапы (Maps)

## Что такое мапа?

**Мапа (map)** в Go — это встроенный тип данных для хранения пар "ключ-значение". Это аналог:
- `HashMap<K,V>` в Java
- `dict` в Python
- `Map` в JavaScript

**Ключевые отличия от Java:**
- В Go `map` — это **встроенный тип**, а не класс (как `HashMap`)
- Синтаксис более лаконичный
- Нет методов `.put()`, `.get()` — работаем напрямую через квадратные скобки
- Нулевое значение мапы — `nil` (как `null` в Java)
- **Nil-мапа нельзя использовать** (panic при записи)

---

## Объявление мапы

### Синтаксис
```go
var m map[KeyType]ValueType
```

Примеры:
```go
var ages map[string]int        // ключ: string, значение: int
var scores map[int]float64     // ключ: int, значение: float64
var cache map[string][]byte    // ключ: string, значение: слайс байтов
```

⚠️ **ВАЖНО:** Объявление через `var` создаёт **nil-мапу**:
```go
var m map[string]int
fmt.Println(m == nil)  // true
// m["key"] = 10       // ❌ PANIC! runtime error
```

**Сравнение с Java:**
```java
// Java
Map<String, Integer> ages = null;  // nil-мapa в Go
ages.put("Alice", 25);  // ❌ NullPointerException

// Правильно в Java
Map<String, Integer> ages = new HashMap<>();
ages.put("Alice", 25);  // ✅
```

---

## Инициализация мапы

### 1. С помощью `make()`
```go
m := make(map[string]int)  // Создаёт пустую мапу, готовую к использованию
m["Alice"] = 25            // ✅ Работает
```

### 2. С начальными значениями (map literal)
```go
ages := map[string]int{
    "Alice": 25,
    "Bob":   30,
    "Charlie": 35,  // Завершающая запятая обязательна!
}
```

### 3. С указанием начальной ёмкости (capacity hint)
```go
m := make(map[string]int, 100)  // Подсказка: ожидаем ~100 элементов
```

**Зачем capacity hint?**
- Аналогично `new HashMap<>(initialCapacity)` в Java
- Уменьшает количество реаллокаций при росте мапы
- Не ограничивает размер (в отличие от слайсов!)

---

## Работа с элементами

### Добавление и обновление
```go
ages := make(map[string]int)
ages["Alice"] = 25      // Добавление
ages["Alice"] = 26      // Обновление (перезапись)
```

**Отличие от Java:**
```java
// Java
ages.put("Alice", 25);   // Добавление
ages.put("Alice", 26);   // Обновление
```

В Go всё через `[]` — проще!

### Чтение значения
```go
age := ages["Alice"]
fmt.Println(age)  // 26
```

⚠️ **Если ключа нет — вернётся zero value:**
```go
age := ages["Zara"]  // Ключа "Zara" нет
fmt.Println(age)     // 0 (zero value для int)
```

**Проблема:** Как отличить "ключа нет" от "ключ есть, значение = 0"?

### Проверка существования ключа (важно!)
```go
age, exists := ages["Alice"]
if exists {
    fmt.Println("Alice's age:", age)
} else {
    fmt.Println("Alice not found")
}
```

**Идиома Go:**
```go
if age, ok := ages["Alice"]; ok {
    fmt.Println("Found:", age)
} else {
    fmt.Println("Not found")
}
```

**Сравнение с Java:**
```java
// Java
if (ages.containsKey("Alice")) {
    int age = ages.get("Alice");
    System.out.println("Found: " + age);
}

// Или через Optional (Java 8+)
Optional.ofNullable(ages.get("Alice"))
    .ifPresent(age -> System.out.println("Found: " + age));
```

---

## Удаление элементов

```go
delete(ages, "Alice")  // Встроенная функция delete()
```

- Если ключа нет — ничего не происходит (не ошибка)
- Возвращает `void` (ничего не возвращает)

**Сравнение с Java:**
```java
ages.remove("Alice");  // Метод .remove()
```

---

## Итерация по мапе

### С помощью `range`
```go
ages := map[string]int{
    "Alice": 25,
    "Bob":   30,
    "Charlie": 35,
}

// Получаем ключ и значение
for name, age := range ages {
    fmt.Printf("%s is %d years old\n", name, age)
}

// Только ключи
for name := range ages {
    fmt.Println(name)
}

// Только значения (игнорируем ключ)
for _, age := range ages {
    fmt.Println(age)
}
```

⚠️ **ВАЖНО: Порядок итерации случайный!**
```go
for name := range ages {
    fmt.Println(name)
}
// Может вывести: Bob, Alice, Charlie
// При следующем запуске: Charlie, Alice, Bob
```

**Это фича, а не баг:**
- В Java `HashMap` тоже не гарантирует порядок
- Но в Go порядок **намеренно рандомизируется** каждый раз
- Чтобы разработчики не полагались на порядок

Если нужен порядок — сортируй ключи:
```go
import "sort"

keys := make([]string, 0, len(ages))
for k := range ages {
    keys = append(keys, k)
}
sort.Strings(keys)

for _, k := range keys {
    fmt.Printf("%s: %d\n", k, ages[k])
}
```

**Сравнение с Java:**
```java
// HashMap — неупорядоченный
Map<String, Integer> ages = new HashMap<>();

// LinkedHashMap — порядок вставки
Map<String, Integer> ages = new LinkedHashMap<>();

// TreeMap — сортированный
Map<String, Integer> ages = new TreeMap<>();
```

В Go только один тип `map` — неупорядоченный.

---

## Длина мапы

```go
fmt.Println(len(ages))  // 3
```

**Нет метода `.size()`** как в Java — только `len()`.

---

## Мапы как reference type

Мапы передаются **по ссылке** (как слайсы):
```go
func addAge(m map[string]int) {
    m["David"] = 40
}

ages := make(map[string]int)
addAge(ages)
fmt.Println(ages["David"])  // 40 — изменения видны снаружи!
```

**Сравнение с Java:**
```java
// Java: объекты тоже передаются "по ссылке" (точнее, передаётся копия ссылки)
void addAge(Map<String, Integer> m) {
    m.put("David", 40);  // Изменяет оригинальную мапу
}
```

---

## Вложенные мапы

Мапа внутри мапы (аналог `Map<String, Map<String, Integer>>` в Java):
```go
scores := map[string]map[string]int{
    "Alice": {
        "Math":    90,
        "Physics": 85,
    },
    "Bob": {
        "Math":    75,
        "Physics": 80,
    },
}

fmt.Println(scores["Alice"]["Math"])  // 90
```

⚠️ **При динамическом создании:**
```go
scores := make(map[string]map[string]int)
scores["Alice"] = make(map[string]int)  // ❗ Нужно создать вложенную мапу!
scores["Alice"]["Math"] = 90
```

---

## Zero value и nil-мапа

```go
var m map[string]int   // nil-мапа
fmt.Println(m == nil)  // true
fmt.Println(len(m))    // 0 (можно узнать длину nil-мапы)

// ✅ Чтение из nil-мапы — ОК (вернёт zero value)
value := m["key"]
fmt.Println(value)  // 0

// ❌ Запись в nil-мапу — PANIC!
// m["key"] = 10  // panic: assignment to entry in nil map
```

**Правило:** Всегда используй `make()` перед записью в мапу.

---

## Сравнение с Java Collections

| Go                           | Java                            |
|------------------------------|---------------------------------|
| `map[string]int`             | `Map<String, Integer>`          |
| `make(map[K]V)`              | `new HashMap<>()`               |
| `m[key] = value`             | `m.put(key, value)`             |
| `value := m[key]`            | `value = m.get(key)`            |
| `value, ok := m[key]`        | `m.containsKey(key)`            |
| `delete(m, key)`             | `m.remove(key)`                 |
| `len(m)`                     | `m.size()`                      |
| `for k, v := range m`        | `for (Map.Entry<K,V> e : m.entrySet())` |
| Нет встроенной синхронизации | `Collections.synchronizedMap()` |

---

## Потокобезопасность

⚠️ **Мапы НЕ потокобезопасны!**

Если несколько горутин одновременно читают и пишут в мапу → **race condition**.

**Решения:**
1. Использовать `sync.Mutex` (изучим позже)
2. Использовать `sync.Map` (специальная потокобезопасная мапа)

**Сравнение с Java:**
```java
// Java
Map<K, V> map = new HashMap<>();  // НЕ потокобезопасна
Map<K, V> map = new ConcurrentHashMap<>();  // Потокобезопасна
```

---

## Ключевые моменты (Summary)

✅ **Мапа** — встроенный тип для хранения пар ключ-значение
✅ **Nil-мапа** — нельзя использовать для записи (только чтение)
✅ **Создание:** `make(map[K]V)` или `map[K]V{...}`
✅ **Проверка ключа:** `value, ok := m[key]`
✅ **Удаление:** `delete(m, key)`
✅ **Итерация:** `for k, v := range m` (порядок случайный!)
✅ **Reference type** — изменения видны при передаче в функции
✅ **НЕ потокобезопасны** по умолчанию

---

## Что дальше?

Теперь ты знаешь основные коллекции в Go:
- **Массивы** — фиксированный размер
- **Слайсы** — динамические массивы
- **Мапы** — ключ-значение

В следующей главе изучим **указатели** (pointers) — как работать с памятью явно (в отличие от Java, где всё через ссылки неявно).

package main

import "fmt"

// TODO: Напиши функцию updateValue, которая принимает указатель на int
// и устанавливает значение по этому адресу в 100

// TODO: Напиши функцию divide, которая принимает два числа (a, b int)
// и два указателя (quotient *int, remainder *int).
// Функция должна записать результат деления a/b в quotient,
// а остаток от деления в remainder.
// Например: divide(10, 3, &q, &r) должно дать q=3, r=1

// TODO: Напиши функцию findMax, которая принимает слайс целых чисел
// и указатель на int. Функция должна записать максимальное значение
// из слайса по указателю. Если слайс пустой, ничего не делай.

// TODO: Напиши функцию increment, которая принимает указатель на int
// и увеличивает значение по этому адресу на 1

// TODO: Напиши функцию swapStrings, которая меняет местами значения
// двух строк используя указатели

// TODO: Напиши функцию createPointer, которая принимает int значение,
// создаёт новую переменную с этим значением, и возвращает указатель на неё.
// Используй оператор & (не new)

// TODO: Напиши функцию safeIncrement, которая принимает указатель на int.
// Если указатель nil, функция ничего не делает.
// Если не nil, увеличивает значение на 1.

func main() {
	fmt.Println("=== Тест 1: updateValue ===")
	// TODO: Создай переменную x = 50
	// TODO: Вызови updateValue(&x)
	// TODO: Выведи x до и после (должно стать 100)

	fmt.Println("\n=== Тест 2: divide ===")
	// TODO: Создай переменные quotient и remainder
	// TODO: Вызови divide(17, 5, &quotient, &remainder)
	// TODO: Выведи результат (должно быть 3 и 2)

	fmt.Println("\n=== Тест 3: findMax ===")
	// TODO: Создай слайс nums := []int{3, 7, 2, 9, 1}
	// TODO: Создай переменную maxVal
	// TODO: Вызови findMax(nums, &maxVal)
	// TODO: Выведи maxVal (должно быть 9)

	// TODO: Протестируй findMax с пустым слайсом
	// TODO: maxEmpty не должен измениться

	fmt.Println("\n=== Тест 4: increment ===")
	// TODO: Создай counter := 10
	// TODO: Вызови increment 3 раза
	// TODO: Выведи counter (должно быть 13)

	fmt.Println("\n=== Тест 5: swapStrings ===")
	// TODO: Создай s1 := "hello", s2 := "world"
	// TODO: Выведи до
	// TODO: Вызови swapStrings(&s1, &s2)
	// TODO: Выведи после (должно поменяться местами)

	fmt.Println("\n=== Тест 6: createPointer ===")
	// TODO: Вызови p := createPointer(42)
	// TODO: Выведи *p (должно быть 42)
	// TODO: Измени *p = 100
	// TODO: Выведи снова (должно быть 100)

	fmt.Println("\n=== Тест 7: safeIncrement ===")
	// TODO: Создай var nilPtr *int
	// TODO: Вызови safeIncrement(nilPtr) — не должно упасть
	// TODO: Создай val := 5
	// TODO: Вызови safeIncrement(&val)
	// TODO: Выведи val (должно быть 6)
}

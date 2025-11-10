package main

import (
	"fmt"
	"slices"
)

// TODO: Напиши функцию updateValue, которая принимает указатель на int
// и устанавливает значение по этому адресу в 100
func updateValue(value *int) {
	*value = 100
}

// TODO: Напиши функцию divide, которая принимает два числа (a, b int)
// и два указателя (quotient *int, remainder *int).
// Функция должна записать результат деления a/b в quotient,
// а остаток от деления в remainder.
// Например: divide(10, 3, &q, &r) должно дать q=3, r=1
func divide(a, b int, quotient, remainder *int) {
	*quotient = a / b
	*remainder = a % b
}

// TODO: Напиши функцию findMax, которая принимает слайс целых чисел
// и указатель на int. Функция должна записать максимальное значение
// из слайса по указателю. Если слайс пустой, ничего не делай.
func findMax(integers []int, max *int) {
	if integers == nil || len(integers) == 0 {
		return
	}

	*max = slices.Max(integers)
}

// TODO: Напиши функцию increment, которая принимает указатель на int
// и увеличивает значение по этому адресу на 1
func increment(value *int) {
	*value++
}

// TODO: Напиши функцию swapStrings, которая меняет местами значения
// двух строк используя указатели
func swapStrings(a, b *string) {
	*a, *b = *b, *a
}

// TODO: Напиши функцию createPointer, которая принимает int значение,
// создаёт новую переменную с этим значением, и возвращает указатель на неё.
// Используй оператор & (не new)
func createPointer(value int) *int {
	result := &value

	return result
}

// TODO: Напиши функцию safeIncrement, которая принимает указатель на int.
// Если указатель nil, функция ничего не делает.
// Если не nil, увеличивает значение на 1.
func safeIncrement(value *int) {
	if value == nil {
		return
	}

	*value++
}

func main() {
	fmt.Println("=== Тест 1: updateValue ===")
	x := 50
	fmt.Println("до:", x)
	updateValue(&x)
	fmt.Println("после:", x) // должно быть 100

	fmt.Println("\n=== Тест 2: divide ===")
	var quotient, remainder int
	divide(17, 5, &quotient, &remainder)
	fmt.Printf("17 / 5 = %d, остаток %d\n", quotient, remainder) // 3, остаток 2

	fmt.Println("\n=== Тест 3: findMax ===")
	nums := []int{3, 7, 2, 9, 1}
	var maxVal int
	findMax(nums, &maxVal)
	fmt.Println("максимум:", maxVal) // 9

	emptySlice := []int{}
	var maxEmpty int
	findMax(emptySlice, &maxEmpty)
	fmt.Println("максимум пустого:", maxEmpty) // 0 (не изменился)

	fmt.Println("\n=== Тест 4: increment ===")
	counter := 10
	increment(&counter)
	increment(&counter)
	increment(&counter)
	fmt.Println("counter после 3 инкрементов:", counter) // 13

	fmt.Println("\n=== Тест 5: swapStrings ===")
	s1, s2 := "hello", "world"
	fmt.Printf("до: s1=%s, s2=%s\n", s1, s2)
	swapStrings(&s1, &s2)
	fmt.Printf("после: s1=%s, s2=%s\n", s1, s2) // s1=world, s2=hello

	fmt.Println("\n=== Тест 6: createPointer ===")
	p := createPointer(42)
	fmt.Println("значение по указателю:", *p) // 42
	*p = 100
	fmt.Println("после изменения:", *p) // 100

	fmt.Println("\n=== Тест 7: safeIncrement ===")
	var nilPtr *int
	safeIncrement(nilPtr) // не должно вызвать panic
	fmt.Println("после safeIncrement(nil): OK")

	val := 5
	safeIncrement(&val)
	fmt.Println("после safeIncrement(&val):", val) // 6
}

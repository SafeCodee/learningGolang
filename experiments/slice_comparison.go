package main

import (
	"fmt"
	"slices"
)

// Вручную: сравнение слайсов
func equalManual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("=== Сравнение слайсов в Go ===\n")

	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	c := []int{1, 2, 4}

	// ❌ a == b не компилируется!
	// fmt.Println(a == b) // invalid operation: a == b (slice can only be compared to nil)

	// ✅ Можно сравнить с nil
	var nilSlice []int
	fmt.Printf("nilSlice == nil: %v\n", nilSlice == nil)
	fmt.Printf("a == nil: %v\n\n", a == nil)

	// ✅ Ручное сравнение
	fmt.Println("=== Ручное сравнение ===")
	fmt.Printf("equalManual(a, b): %v\n", equalManual(a, b))   // true
	fmt.Printf("equalManual(a, c): %v\n\n", equalManual(a, c)) // false

	// ✅ Стандартная библиотека (Go 1.21+)
	fmt.Println("=== slices.Equal (Go 1.21+) ===")
	fmt.Printf("slices.Equal(a, b): %v\n", slices.Equal(a, b)) // true
	fmt.Printf("slices.Equal(a, c): %v\n", slices.Equal(a, c)) // false

	// Сравнение с разной длиной
	d := []int{1, 2, 3, 4}
	fmt.Printf("slices.Equal(a, d): %v\n\n", slices.Equal(a, d)) // false

	// Дополнительно: поиск элементов
	fmt.Println("=== Поиск элементов в слайсах ===")
	nums := []int{10, 20, 30, 40, 50}

	// Go 1.21+: slices.Contains
	fmt.Printf("slices.Contains(nums, 30): %v\n", slices.Contains(nums, 30)) // true
	fmt.Printf("slices.Contains(nums, 99): %v\n", slices.Contains(nums, 99)) // false

	// Индекс элемента
	fmt.Printf("slices.Index(nums, 30): %d\n", slices.Index(nums, 30)) // 2
	fmt.Printf("slices.Index(nums, 99): %d\n", slices.Index(nums, 99)) // -1 (не найден)
}

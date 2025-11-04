package main

import "fmt"

func main() {
	fmt.Println("=== Подводный камень: append перезаписывает общий массив ===\n")

	original := []int{1, 2, 3, 4, 5}
	slice1 := original[:3]  // [1 2 3], cap=5
	slice2 := original[2:5] // [3 4 5], cap=3

	fmt.Printf("До append:\n")
	fmt.Printf("original: %v\n", original)
	fmt.Printf("slice1:   %v, cap=%d\n", slice1, cap(slice1))
	fmt.Printf("slice2:   %v, cap=%d\n\n", slice2, cap(slice2))

	// Добавляем в slice1 (влезает в capacity!)
	slice1 = append(slice1, 999)

	fmt.Printf("После append(slice1, 999):\n")
	fmt.Printf("original: %v — изменился!\n", original)
	fmt.Printf("slice1:   %v\n", slice1)
	fmt.Printf("slice2:   %v — тоже изменился!\n\n", slice2)

	fmt.Println("=== Решение: использовать full slice expression [low:high:max] ===\n")

	original2 := []int{1, 2, 3, 4, 5}
	safe1 := original2[:3:3] // capacity ограничен до 3
	safe2 := original2[2:]

	fmt.Printf("До append:\n")
	fmt.Printf("original2: %v\n", original2)
	fmt.Printf("safe1:     %v, cap=%d\n", safe1, cap(safe1))
	fmt.Printf("safe2:     %v, cap=%d\n\n", safe2, cap(safe2))

	safe1 = append(safe1, 999) // не влезает в capacity -> создаётся новый массив

	fmt.Printf("После append(safe1, 999):\n")
	fmt.Printf("original2: %v — НЕ изменился!\n", original2)
	fmt.Printf("safe1:     %v — в новом массиве\n", safe1)
	fmt.Printf("safe2:     %v — НЕ изменился!\n", safe2)
}

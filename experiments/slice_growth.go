package main

import "fmt"

func main() {
	fmt.Println("=== Рост capacity в Go ===")
	s := make([]int, 0, 1) // capacity = 1

	for i := 0; i < 20; i++ {
		prevCap := cap(s)
		s = append(s, i)
		newCap := cap(s)

		if prevCap != newCap {
			fmt.Printf("len=%2d: capacity вырос %d -> %d (коэффициент: %.2fx)\n",
				len(s), prevCap, newCap, float64(newCap)/float64(prevCap))
		}
	}

	fmt.Println("\n=== Большие слайсы (> 256) ===")
	big := make([]int, 0, 256)
	for i := 0; i < 1000; i++ {
		prevCap := cap(big)
		big = append(big, i)
		newCap := cap(big)

		if prevCap != newCap && prevCap >= 256 {
			fmt.Printf("capacity вырос %d -> %d (коэффициент: %.2fx)\n",
				prevCap, newCap, float64(newCap)/float64(prevCap))
		}
	}
}

package main

import (
	"fmt"
	"time"
)

func pow(n uint64, m uint64) uint64 {
	result := uint64(1)
	for ; m > 0; m /= 2 {
		if m%2 == 1 {
			result *= n
		}
		n *= n
	}
	return result
}

func somesum(n uint64, m uint64) uint64 {
	sum := uint64(0)
	for i := uint64(0); i <= n; i++ {
		sum += pow(i, m)
	}
	return sum
}

func main() {
	start := time.Now()
	c := somesum(60000, 3) // Код для измерения времени исполнения
	duration := time.Since(start)
	fmt.Println(duration)
	fmt.Printf("\n%d", c)
}

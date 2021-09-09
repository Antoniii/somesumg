package main

import (
	"fmt"
	"math"
	"time"
)

func somesum(n uint64, m uint64) uint64 {
	sum := 0.0
	var i uint64
	for i = 0; i <= n; i++ {
        sum += math.Pow(float64(i), float64(m))
    }
    return uint64(sum)
}

func main() {
	start := time.Now()
	c := somesum(60000, 3) // Код для измерения времени исполнения
	duration := time.Since(start)
	fmt.Println(duration)
	fmt.Printf("\n%d", c)
}

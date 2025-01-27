package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(MinInt(10, 15))
}

func MinInt(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

package main

import (
	"fmt"
	"math"
)

func main() {
	x := 1534236469
	fmt.Println((2 ^ 31) - 1)
	fmt.Println(reverse(x))
}

func reverse(x int) int {
	if x < 10 && x > -10 {
		return x
	}

	isNegative := false

	if x < 0 {
		isNegative = true
		x = 0 - x
	}

	step := 10
	result := 0

	for x > 0 {
		remainder := x % 10
		result = result*step + remainder
		x /= 10
	}

	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}

	if isNegative {
		return 0 - result
	}
	return result
}

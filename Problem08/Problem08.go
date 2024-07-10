package main

import (
	"fmt"
	"math"
)

func main() {
	s := " ++1"
	fmt.Println(myAtoi(s))
}

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	i := 0
	mark := 0

	// Check if any character is space, +, -
	for i < len(s) && (s[i] == 32 || s[i] == 43 || s[i] == 45) {
		// If has +/- before and current character is space
		if mark != 0 && s[i] == 32 {
			return 0
		}

		// If has +/- before and current character is +/-
		if mark != 0 && (s[i] == 43 || s[i] == 45) {
			return 0
		}

		// If has -
		if s[i] == 45 {
			mark = 45
		}

		// If has +
		if s[i] == 43 {
			mark = 43
		}
		i++
	}

	// Check if not have any digit
	if i >= len(s) {
		return 0
	}

	// Check if has any non digit at beginning
	if s[i] > 57 || s[i] < 48 {
		return 0
	}

	step := 10
	result := 0

	// Build number
	for i := i; i < len(s); i++ {
		if s[i] <= 57 && s[i] >= 48 {
			result = result*step + int(s[i]) - 48
		} else {
			break
		}

		// Check if out of integer bound
		if result > math.MaxInt32 {
			if mark == 45 {
				return math.MinInt32
			} else {
				return math.MaxInt32
			}
		}
	}

	if mark == 45 {
		return 0 - result
	} else {
		return result
	}
}

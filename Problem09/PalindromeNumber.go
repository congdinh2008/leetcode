package main

import "fmt"

func main() {
	x := 121
	result := IsPalindrome(x)
	fmt.Println(result)
}

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x == 0 {
		return true
	}

	if x%10 == 0 {
		return false
	}

	temp := x
	reversed := 0

	for temp > 0 {
		reversed = reversed*10 + temp%10
		temp /= 10
		if temp/10 == reversed || temp == reversed {
			return true
		}
	}

	return x == reversed
}

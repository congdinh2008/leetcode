package main

import (
	"fmt"
)

func main() {
	s := "abceeabdcxbb"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	myMap := map[rune]int{}
	maxLength := 0
	left := 0

	for right, char := range s {
		if index, found := myMap[char]; found && index >= left {
			fmt.Println("index: ", index, "left: ", left)
			left = index + 1
		}
		myMap[char] = right
		currentLength := right - left + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}

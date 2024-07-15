package main

import (
	"fmt"
)

func main() {
	digits := "234"

	for _, item := range letterCombinations(digits) {
		fmt.Println(item)
	}
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	mapping := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	var result []string
	backtrack(&result, mapping, digits, "", 0)
	return result
}

func backtrack(result *[]string, mapping map[string]string, digits, s string, i int) {
	if i == len(digits) {
		*result = append(*result, s)
		return
	}

	for _, item := range mapping[string(digits[i])] {
		backtrack(result, mapping, digits, s+string(item), i+1)
	}
}

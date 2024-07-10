package main

import (
	"fmt"
	"strings"
)

func main() {
	strs := []string{"ab", "a"}
	fmt.Println(longestCommonPrefix02(strs))
}

func longestCommonPrefix(strs []string) string {
	var result strings.Builder
	result.WriteString("")

	i := 0
	isBreak := false

	if strs[0] == "" {
		return result.String()
	}

	for {
		if i >= len(strs[0]) {
			break
		}

		currentChar := strs[0][i]

		for _, item := range strs {
			if i >= len(item) || item[i] != currentChar || item == "" {
				isBreak = true
				break
			}
		}
		if isBreak {
			break
		} else {
			result.WriteByte(currentChar)
			i++
		}
	}

	return result.String()
}

func longestCommonPrefix02(strs []string) string {
	result := strs[0]

	for _, item := range strs {
		for len(item) < len(result) || item[:len(result)] != result {
			result = result[:len(result)-1]
			if result == "" {
				return result
			}
		}
	}
	return result
}

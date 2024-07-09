package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abadba"
	fmt.Println(longestPalindromeDP02(s))
}

func longestPalindromeDP02(s string) string {
	n := len(s)
	startIndex := 0
	maxlen := 0
	table := make([][]bool, n)
	for i := range table {
		table[i] = make([]bool, n)
	}

	for g := 0; g < n; g++ {
		for i, j := 0, g; j < n; i, j = i+1, j+1 {
			solve(table, i, j, s)
			if table[i][j] {
				if j-i+1 > maxlen {
					startIndex = i
					maxlen = j - i + 1
				}
			}
		}
	}
	return s[startIndex : startIndex+maxlen]
}

func solve(table [][]bool, i, j int, s string) bool {
	if i == j {
		table[i][j] = true
		return true
	}
	if j-i == 1 {
		if s[i] == s[j] {
			table[i][j] = true
			return true
		} else {
			table[i][j] = false
			return false
		}
	}
	if s[i] == s[j] && table[i+1][j-1] {
		table[i][j] = true
		return true
	} else {
		table[i][j] = false
		return false
	}
}

func longestPalindromeDP(s string) string {
	length := len(s)
	table := make([][]bool, length)

	for i := range s {
		table[i] = make([]bool, length)
		table[i][i] = true
	}

	maxLength := 1
	start := 0
	for i := 3; i <= length; i++ {
		for j := 0; j < length-i+1; j++ {
			k := i + j - 1
			if table[j+1][k-1] && s[j] == s[k] {
				table[j][k] = true
				if i > maxLength {
					start = j
					maxLength = i
				}
			}
		}
	}
	return s[start:maxLength]
}

func longestPalindrome02(s string) string {
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		oddPalindrome := expandFromCenter(s, i, i)
		evenPalindrome := expandFromCenter(s, i, i+1)

		if len(oddPalindrome) > result.Len() {
			result.Reset()
			result.WriteString(oddPalindrome)
		}

		if len(evenPalindrome) > result.Len() {
			result.Reset()
			result.WriteString(evenPalindrome)
		}
	}
	return result.String()
}

func expandFromCenter(s string, left int, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left -= 1
		right += 1
	}

	return s[left+1 : right]
}

func longestPalindrome(s string) string {
	var result strings.Builder

	for i := 0; i < len(s); i++ {
		temp := s[i:]
		tempLen := len(temp)
		for tempLen > 0 {
			if isPalindrome(temp) {
				if len(temp) > result.Len() {
					result.Reset()
					result.WriteString(temp)
				}
				break
			} else {
				tempLen--
				temp = temp[0:tempLen]
			}
		}
	}
	return result.String()
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

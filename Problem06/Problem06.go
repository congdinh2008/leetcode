package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "PAYPALISHIRING"
	numRows := 3

	fmt.Println(convert(s, numRows))
}

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	var result strings.Builder
	// The term is the number of characters in one zigzag cycle
	term := 2*numRows - 2

	for i := 0; i < numRows; i++ {
		for j := i; j < len(s); j += term {
			// Always add the current character in the row
			result.WriteByte(s[j])
			// For non-first and non-last rows, add the character that appears diagonally in the zigzag
			if i != 0 && i != numRows-1 {
				diagIndex := j + term - 2*i
				if diagIndex < len(s) {
					result.WriteByte(s[diagIndex])
				}
			}
		}
	}

	return result.String()
}

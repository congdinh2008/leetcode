package main

func main() {
	s := "()[]{}"
	println(isValid(s))
}

func isValid02(s string) bool {
	str := ""
	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			str += string(char)
		} else {
			if len(str) == 0 {
				return false
			}

			lastChar := str[len(str)-1]
			if (char == ')' && lastChar == '(') || (char == ']' && lastChar == '[') || (char == '}' && lastChar == '{') {
				str = str[:len(str)-1]
			} else {
				return false
			}
		}
	}
	return len(str) == 0
}

func isValid(s string) bool {
	var stack Stack
	pairs := map[rune]string{
		'(': ")",
		'[': "]",
		'{': "}",
	}

	for _, char := range s {
		if closing, exists := pairs[char]; exists {
			stack.Push(closing)
		} else {
			if len(stack) == 0 || stack.Pop() != string(char) {
				return false
			}
		}
	}
	return len(stack) == 0
}

type Stack []string

func (stack *Stack) Push(str string) {
	*stack = append(*stack, str)
}

func (stack *Stack) Pop() string {
	if len(*stack) == 0 {
		return ""
	}

	index := len(*stack) - 1
	element := (*stack)[index]
	*stack = (*stack)[:index]
	return element
}

func (stack *Stack) Peek() string {
	if len(*stack) == 0 {
		return ""
	}

	index := len(*stack) - 1
	element := (*stack)[index]
	return element
}

package main

import "fmt"

func main() {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	l3 := addTwoNumbers(l1, l2)

	for l3 == nil {
		fmt.Println(l3.Val)
		l3 = l3.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	current := &ListNode{0, nil}
	result := current
	sum := 0

	for l1 != nil || l2 != nil {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		current.Val = sum % 10
		sum /= 10

		if l1 != nil || l2 != nil {
			current.Next = &ListNode{0, nil}
			current = current.Next
		}
	}

	if sum > 0 {
		current.Next = &ListNode{sum, nil}
		current = current.Next
	}

	return result
}

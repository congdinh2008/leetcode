package main

func main() {
	// test data
	// head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}
	n := 2
	head = removeNthFromEnd(head, n)

	// print result
	for head != nil {
		println(head.Val)
		head = head.Next
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		head = nil
		return head
	}

	right := 0
	left := head

	current := head
	for current.Next != nil {
		right++
		if right-n > 0 {
			left = left.Next
		}
		current = current.Next
	}

	if right+1 == n {
		head = head.Next
		return head
	}

	left.Next = left.Next.Next

	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}

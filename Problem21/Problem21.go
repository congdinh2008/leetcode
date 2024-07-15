package main

func main() {
	list1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	result := mergeTwoLists(list1, list2)

	for result != nil {
		println(result.Val)
		result = result.Next
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
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	current := &ListNode{0, nil}
	result := current

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Val = list1.Val
			list1 = list1.Next
		} else {
			current.Val = list2.Val
			list2 = list2.Next
		}

		if list1 != nil || list2 != nil {
			current.Next = &ListNode{0, nil}
			current = current.Next
		}
	}

	for list1 != nil {
		current.Val = list1.Val
		list1 = list1.Next

		if list1 != nil {
			current.Next = &ListNode{0, nil}
			current = current.Next
		}
	}

	for list2 != nil {
		current.Val = list2.Val
		list2 = list2.Next

		if list2 != nil {
			current.Next = &ListNode{0, nil}
			current = current.Next
		}
	}

	return result
}

func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists2(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists2(list1, list2.Next)
		return list2
	}
}

func mergeTwoLists3(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	current := result

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	}

	if list2 != nil {
		current.Next = list2
	}

	return result.Next
}

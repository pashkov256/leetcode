package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func createList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &ListNode{Val: values[0]}
	current := head

	for _, val := range values[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}

	return head
}

func printValueList(list *ListNode) {
	curr := list

	for curr != nil {
		fmt.Println(curr.Val)
		curr = curr.Next
	}
}

func deleteDuplicates(head *ListNode) *ListNode {
	curr := head

	for curr != nil && curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}

	}
	return head
}

func main() {
	list := createList([]int{1, 1, 2, 3, 3, 4})

	printValueList(deleteDuplicates(list))

}

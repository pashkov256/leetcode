package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func createLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
	}
	return head
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := head

	for curr != nil {
		fmt.Println(curr)
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	
	return prev
}

func main() {
	head := createLinkedList([]int{1, 2})
	fmt.Print("Original list: ", head)
	reverseList(head)
}

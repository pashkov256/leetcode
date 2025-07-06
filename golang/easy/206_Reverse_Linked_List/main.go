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
	newList := &ListNode{}
	for i := 0; i < 10; i++ {

		fmt.Println(i)
	}
	return newList
}

func main() {
	head := createLinkedList([]int{1, 2})
	fmt.Print("Original list: ", head)
	reverseList(head)
}

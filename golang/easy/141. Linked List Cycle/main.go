package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	fast,slow := head,head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}//второе решение

// func hasCycle(head *ListNode) bool {
// 	listMap := make(map[*ListNode]struct{},0) 

//     if head != nil {

//         for head.Next != nil {

// 			if _,ok := listMap[head];!ok {
// 				listMap[head] = struct{}{}
// 			} else {
// 				return true
// 			}

// 			head = head.Next
// 		}
//     }

// 	return false
// }//ПЕРВОЕ РЕШЕНИЕ

func main() {

	node1 := &ListNode{Val: 3}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 0}
	node4 := &ListNode{Val: -4}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2 

	// Голова списка
	headWithCycle := node1
	fmt.Println(hasCycle(headWithCycle))
}
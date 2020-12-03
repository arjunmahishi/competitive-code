// Problem link: https://leetcode.com/problems/add-two-numbers

package leetcode

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

type List struct {
	total int
	head  *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return nil
}

func llToStack(ll *ListNode) []int {
	stack := []int{}
	curr := ll
	for curr != nil {
		fmt.Println(curr.Val, "value")
		stack = append(stack, curr.Val)
		fmt.Println(curr.Next, "value2")
		curr = curr.Next
	}
	return stack
}

func stackToLL(stack []int) *List {
	newList := &List{}
	for i := len(stack) - 1; i >= 0; i-- {
		newList.add(stack[i])
	}
	return newList
}

func (l *List) add(i int) {
	newNode := &ListNode{
		Val:  i,
		Next: nil,
	}
	curr := l.head
	for curr != nil {
		fmt.Println(curr.Val, "val")
		curr = curr.Next
	}
	curr = newNode
	l.total++
}

// SPDX-License-Identifier: MIT
// Copyright (c) 2019 Gennady Trafimenkov

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// insert inserts a new node into the head of the list
// and returns the new head.
func insert(n *ListNode, list *ListNode) *ListNode {
	if list == nil {
		return n
	}
	n.Next = list
	return n
}

func printList(list *ListNode) {
	fmt.Printf("%p", list)
	for list.Next != nil {
		fmt.Printf(" -> %p", list.Next)
		list = list.Next
	}
	fmt.Println("")
}

// hasCycleM checks if a list has a cycle.
// The solution uses extra memory: pointers to all visited nodes are kept in a hash.
func hasCycleM(head *ListNode) bool {
	visitedNodes := map[*ListNode]bool{}
	for head != nil {
		v := visitedNodes[head]
		if v {
			return true
		}
		visitedNodes[head] = true
		head = head.Next
	}
	return false
}

// hasCycleC checks if a list has a cycle.
//
// The solution doesn't use extra memory, but more computationaly intensive.
// Computantional complexity: O(n)
//
// Algorithm:
//   1. Select the first node of the list as the checkpoint.
//   2. Make N steps.  If the checkpoint is encountered, the list has a cycle.  If the of the list is reached,
//      the list doesn't have a cycle.
//   3. Chose the current item of the list to be next checkpoint.  Double N.  Go to 1.  Eventually either
//      the end of the list will be reached or the checkpoint encountered.
func hasCycleC(head *ListNode) bool {
	if head == nil {
		return false
	}

	n := 10
	checkpoint := head
	current := head
	for {
		for i := n; i > 0; i-- {
			current = current.Next
			if current == nil {
				return false
			}
			if current == checkpoint {
				return true
			}
		}
		checkpoint = current
		n = n * 2
	}
}

func hasCycle(list *ListNode) {
	fmt.Printf("m:\t%v\n", hasCycleM(list))
	fmt.Printf("c:\t%v\n", hasCycleC(list))
	fmt.Println("------------------")
}

func main() {

	{
		hasCycle(nil)
	}

	{
		n1 := &ListNode{}
		list := insert(n1, nil)
		hasCycle(list)
	}

	{
		n1 := &ListNode{}
		list := insert(n1, nil)

		// making a cycle
		list.Next = list

		hasCycle(list)
	}

	{
		n1 := &ListNode{}
		n2 := &ListNode{}
		n3 := &ListNode{}
		n4 := &ListNode{}
		list := insert(n1, insert(n2, insert(n3, insert(n4, nil))))
		hasCycle(list)
	}

	{
		n1 := &ListNode{}
		n2 := &ListNode{}
		n3 := &ListNode{}
		n4 := &ListNode{}
		list := insert(n1, insert(n2, insert(n3, insert(n4, nil))))

		// making a cycle
		n4.Next = n2

		hasCycle(list)
	}
}

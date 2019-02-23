// SPDX-License-Identifier: MIT
// Copyright (c) 2019 Gennady Trafimenkov

package main

import "fmt"

type node struct {
	next *node
}

// insert inserts a new node into the head of the list
// and returns the new head.
func insert(n *node, list *node) *node {
	if list == nil {
		return n
	}
	n.next = list
	return n
}

func printList(list *node) {
	fmt.Printf("%p", list)
	for list.next != nil {
		fmt.Printf(" -> %p", list.next)
		list = list.next
	}
	fmt.Println("")
}

// hasCycleM checks if a list has a cycle.
// The solution uses extra memory: pointers to all visited nodes are kept in a hash.
func hasCycleM(list *node) bool {
	visitedNodes := map[*node]bool{}
	for list != nil {
		v := visitedNodes[list]
		if v {
			return true
		}
		visitedNodes[list] = true
		list = list.next
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
func hasCycleC(list *node) bool {
	if list == nil {
		return false
	}

	n := 10
	checkpoint := list
	current := list
	for {
		for i := n; i > 0; i-- {
			current = current.next
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

func hasCycle(list *node) {
	fmt.Printf("m:\t%v\n", hasCycleM(list))
	fmt.Printf("c:\t%v\n", hasCycleC(list))
	fmt.Println("------------------")
}

func main() {

	{
		hasCycle(nil)
	}

	{
		n1 := &node{}
		list := insert(n1, nil)
		hasCycle(list)
	}

	{
		n1 := &node{}
		list := insert(n1, nil)

		// making a cycle
		list.next = list

		hasCycle(list)
	}

	{
		n1 := &node{}
		n2 := &node{}
		n3 := &node{}
		n4 := &node{}
		list := insert(n1, insert(n2, insert(n3, insert(n4, nil))))
		hasCycle(list)
	}

	{
		n1 := &node{}
		n2 := &node{}
		n3 := &node{}
		n4 := &node{}
		list := insert(n1, insert(n2, insert(n3, insert(n4, nil))))

		// making a cycle
		n4.next = n2

		hasCycle(list)
	}
}

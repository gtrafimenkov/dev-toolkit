// SPDX-License-Identifier: MIT
// Copyright (c) 2019 Gennady Trafimenkov

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (node TreeNode) String() string {
	return fmt.Sprintf("(%v %s %s)", node.Val, node.Left, node.Right)
}

func (node *TreeNode) add(value int) {
	storeValue := func(leaf **TreeNode, value int) {
		if *leaf != nil {
			(*leaf).add(value)
		} else {
			*leaf = &TreeNode{
				Val:   value,
				Left:  nil,
				Right: nil,
			}
		}
	}
	if node.Val < value {
		storeValue(&node.Right, value)
	}

	if node.Val > value {
		storeValue(&node.Left, value)
	}
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	return &TreeNode{
		Val:   root.Val,
		Left:  invertTree(root.Right),
		Right: invertTree(root.Left),
	}
}

func main() {
	root := TreeNode{Val: 4}
	root.add(2)
	root.add(7)
	root.add(1)
	root.add(3)
	root.add(6)
	root.add(9)
	fmt.Println(root)

	invertedTree := invertTree(&root)
	fmt.Println(invertedTree.String())
}

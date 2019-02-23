// SPDX-License-Identifier: MIT
// Copyright (c) 2019 Gennady Trafimenkov

package main

import "fmt"

func maxArea(height []int) int {
	size := len(height)
	max := 0
	for i := 0; i < size-1; i++ {
		leftHeight := height[i]
		if leftHeight > 0 {
			for j := i + 1; j < size; j++ {
				height := height[j]
				if height > 0 {
					if height > leftHeight {
						height = leftHeight
					}
					v := height * (j - i)
					if v > max {
						max = v
					}
				}
			}
		}
	}
	return max
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

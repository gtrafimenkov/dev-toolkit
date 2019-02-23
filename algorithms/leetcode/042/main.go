// SPDX-License-Identifier: MIT
// Copyright (c) 2019 Gennady Trafimenkov

package main

import "fmt"

func trap(height []int) int {
	maxHeight := 0
	for _, h := range height {
		if h > maxHeight {
			maxHeight = h
		}
	}

	acc := 0

	for level := 1; level <= maxHeight; level++ {
		inBasin := false
		waterStretch := 0
		for _, h := range height {
			if h < level {
				if inBasin {
					waterStretch++
				}
			} else {
				// end of basic and possible start of a new one
				if inBasin {
					acc += waterStretch
					waterStretch = 0
				}
				inBasin = true
			}
		}
	}

	return acc
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

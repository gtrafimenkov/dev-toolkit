// SPDX-License-Identifier: MIT
// Copyright (c) 2019 Gennady Trafimenkov

package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	r := 0
	results := make([]int, len(nums)-k+1)
	length := len(nums)

	i := 1
	var prevMax int

	// calculating the first max
	{
		max := nums[0]
		for ; i < k; i++ {
			val := nums[i]
			if val > max {
				max = val
			}
		}
		results[r] = max
		r++
		prevMax = max
	}

	for ; i < length; i++ {
		leaving := nums[i-k]
		if leaving < prevMax {
			// no need to recalculate max of the rest of the window
			if nums[i] > prevMax {
				prevMax = nums[i]
			}
		} else {
			max := nums[i-k+1]
			for j := i - k + 2; j <= i; j++ {
				val := nums[j]
				if val > max {
					max = val
				}
			}
			prevMax = max
		}
		results[r] = prevMax
		r++
	}
	return results
}

func maxSlidingWindowSimple(nums []int, k int) []int {
	results := []int{}
	if len(nums) > 0 {
		end := len(nums) - k
		for start := 0; start <= end; start++ {
			max := nums[start]
			windowEnd := start + k
			for j := start + 1; j < windowEnd; j++ {
				val := nums[j]
				if val > max {
					max = val
				}
			}
			results = append(results, max)
		}
	}
	return results
}

func maxSlidingWindowSimple2(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	r := 0
	results := make([]int, len(nums)-k+1)

	end := len(nums) - k
	for start := 0; start <= end; start++ {
		max := nums[start]
		windowEnd := start + k
		for j := start + 1; j < windowEnd; j++ {
			val := nums[j]
			if val > max {
				max = val
			}
		}
		results[r] = max
		r++
	}
	return results
}

func main() {
	fmt.Println(maxSlidingWindowSimple2([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(maxSlidingWindowSimple2([]int{}, 0))
}

// SPDX-License-Identifier: MIT
// Copyright (c) 2019 Gennady Trafimenkov

// https://leetcode.com/problems/palindrome-number/

package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	top := 10
	for top < x {
		top *= 10
	}
	top /= 10

	for top > 0 {
		if x/top != x%10 {
			return false
		}
		x = (x % top) / 10
		top /= 100
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(-10))
	fmt.Println(isPalindrome(0))
	fmt.Println(isPalindrome(1))
	fmt.Println(isPalindrome(11))
	fmt.Println(isPalindrome(111))
	fmt.Println(isPalindrome(1121))
	fmt.Println(isPalindrome(555551155555))
	fmt.Println(isPalindrome(5555511555550))
	fmt.Println(isPalindrome(1000021))
}

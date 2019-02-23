// SPDX-License-Identifier: MIT
// Copyright (c) 2019 Gennady Trafimenkov

package main

import "fmt"

func isBalanced(s string) bool {
	mapping := map[rune]rune{
		'[': ']',
		'{': '}',
		'(': ')',
	}
	openBrackets := []rune{}
	for _, r := range s {
		if r == '[' || r == '{' || r == '(' {
			openBrackets = append(openBrackets, r)
			continue
		}
		if r == ']' || r == '}' || r == ')' {
			if len(openBrackets) == 0 {
				return false
			}
			lastOpen := openBrackets[len(openBrackets)-1]
			if mapping[lastOpen] != r {
				return false
			}
			openBrackets = openBrackets[0 : len(openBrackets)-1]
			continue
		}
	}
	return len(openBrackets) == 0
}

func main() {
	tests := []string{
		"{}()[{}]",
		"[({})]",
		"({[]})",
		"[({)}]",
		"({[})",
		"()}[]",
		"(",
		")",
		"",
		"foo",
		"\n",
	}
	for _, s := range tests {
		fmt.Printf("%s\t%v\n", s, isBalanced(s))
	}
}

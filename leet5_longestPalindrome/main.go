package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestPalindrome("cbbd"))
}

func longestPalindrome(s string) string {
	res := ""
	l := len(s)
	if l == 0 || l == 1 {
		return s
	}

	var extend func(left, right int)
	extend = func(left, right int) {
		for left >= 0 && right < len(s) {
			if s[left] == s[right] {
				left--
				right++
			} else {
				break
			}
		}
		if right-left-1 > len(res) {
			res = s[left+1 : right]
		}
		fmt.Println(res)
	}
	for i := 0; i < len(s); i++ {
		extend(i, i+1)
		extend(i, i)
	}
	return res
}

package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestPalindrome("abcddcba"))
}

func longestPalindrome(s string) string {
	res := ""
	l := len(s)
	var extend func(int, int)
	extend = func(left int, right int) {
		for left >= 0 && right < l {
			if s[left] == s[right] {
				if right-left > len(res) {
					res = s[left : right+1]
				}
				left--
				right++
			} else {
				return
			}
		}
	}

	for i := 0; i < len(s); i++ {
		extend(i, i)
		extend(i, i+1)
	}
	return res

}

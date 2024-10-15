package main

import (
	"fmt"
	"math"
)

// 滑动窗口
// 当移动right 扩大窗口，即加入字符时， 应该更新那些数据？
// 什么条件下， 窗口应该暂停扩大， 开始移动left缩小窗口
// 当移动left缩小窗口， 即移出字符时， 应该更新那些数据
// 我们要的结果应该在扩大窗口时还是缩小窗口时进行更新

func main() {
	need := make(map[byte]int)
	need['s'] = 1
	need['a'] = 2
	fmt.Println(math.MaxInt32)
}

func minWindow(s, t string) string {
	need, windows := make(map[byte]int), make(map[byte]int)
	for i := range t {
		need[t[i]]++
	}

	left, right := 0, 0
	start, length := 0, math.MaxInt32
	valid := 0
	for right < len(s) {
		c := s[right]
		right++
		if _, ok := need[c]; ok {
			windows[c]++
			if windows[c] == need[c] {
				valid++
			}
		}

		for valid == len(need) {
			if right-left < length {
				start = left
				length = right - left
			}
			d := s[left]
			left++
			if _, ok := need[d]; ok {
				if windows[d] == need[d] {
					valid--
				}
				windows[d]--
			}
		}

	}
	if length == math.MaxInt32 {
		return ""
	}
	return s[start : start+length]

}

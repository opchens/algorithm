package main

import "fmt"

// 滑动窗口
// 当移动right 扩大窗口，即加入字符时， 应该更新那些数据？
// 什么条件下， 窗口应该暂停扩大， 开始移动left缩小窗口
// 当移动left缩小窗口， 即移出字符时， 应该更新那些数据
// 我们要的结果应该在扩大窗口时还是缩小窗口时进行更新

func main() {
	fmt.Println(minWindows("s", "sade"))
}

func minWindows(s, t string) string {
	destCount := [128]int{}
	for _, v := range t {
		destCount[v]++
	}

	curCount := [128]int{}
	deffer := len(t)

	minIdx, minLen := -1, len(s)+1
	left := 0
	for right := 0; right < len(s); right++ {
		if destCount[s[right]] < 1 {
			continue
		}
		curCount[s[right]]++

		if curCount[s[right]] <= destCount[s[right]] {
			deffer--
		}
		if deffer == 0 {
			if destCount[s[left]] == 0 {
				left++
				continue
			}
			curLen := right - left + 1
			if curLen < minLen {
				minLen = curLen
				minIdx = left
			}
			if curCount[s[left]] < destCount[s[left]] {
				deffer++
			}
			left++
		}
	}
	if minIdx == -1 {
		return ""
	}

	return s[minIdx : minIdx+minLen]

}

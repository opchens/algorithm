package main

func main() {
	rotate([]int{1, 2, 3, 4}, 2)
}

func rotate(nums []int, k int) {
	l := len(nums)
	k = k % l
	nums = append(nums[l-k:], nums[:l-k]...)
	// leetcode 用 copy
}

package main

func sortedSquares(nums []int) []int {
	left := 0
	l, right := len(nums)-1, len(nums)-1
	res := make([]int, l+1)

	for left <= right {
		if nums[left]*nums[left] > nums[right]*nums[right] {
			res[l] = nums[left] * nums[left]
			left++
		} else {
			res[l] = nums[right] * nums[right]
			right--
		}
		l--
	}
	return res
}

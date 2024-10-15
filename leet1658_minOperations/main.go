package main

func minOperations(nums []int, x int) int {
	n := len(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}

	target := sum - x

	windowsSum := 0
	left, right := 0, 0
	maxLen := -1

	for right < n {
		windowsSum = nums[right]
		right++

		for windowsSum > target && left < right {
			windowsSum -= nums[left]
			left++
		}
		if windowsSum == target {
			if maxLen == -1 || right-left > maxLen {
				maxLen = right - left
			}
		}
	}

	if maxLen == -1 {
		return maxLen
	}
	return n - maxLen
}

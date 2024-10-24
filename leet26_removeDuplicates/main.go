package main

func main() {}

func removeDuplicates(nums []int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[slow] == nums[fast] {
			fast += 1
		} else {
			slow += 1
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}

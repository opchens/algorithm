package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 9}
	fmt.Println(Search(nums, 6))
}

func Search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (right-left)/2 + left
		if target == nums[mid] {
			return mid
		}
		if target < nums[mid] {
			right = mid - 1
		}
		if target > nums[mid] {
			left = mid + 1
		}
	}
	return -1
}

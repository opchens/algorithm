package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(fourSum([]int{-1, 0, 1, 2, -1, -4}, 0))
}

func fourSum(nums []int, target int) [][]int {

	var res [][]int

	sort.Ints(nums)
	l := len(nums) - 1

	for i := 0; i < l; i++ {
		tsum := threeSum(nums, i+1, l, target-nums[i])
		for _, t := range tsum {
			t = append(t, nums[i])
			res = append(res, t)
		}
		for i < l && nums[i] == nums[i+1] {
			i++
		}
		fmt.Println(res)
	}
	return res
}

func threeSum(nums []int, start, end, target int) [][]int {

	var sum [][]int
	for i := start; i < end; i++ {
		tsum := twoSum(nums, i+1, end, target-nums[i])
		for _, t := range tsum {
			t = append(t, nums[i])
			sum = append(sum, t)
		}
		for i < end && nums[i] == nums[i+1] {
			i++
		}
	}
	return sum

}

func twoSum(nums []int, start, end, target int) [][]int {
	var sum [][]int
	for start < end {
		if nums[start]+nums[end] < target {
			start++
		} else if nums[start]+nums[end] > target {
			end--
		} else if nums[start]+nums[end] == target {
			sum = append(sum, []int{nums[start], nums[end]})
			for start < end && nums[start] == nums[start+1] {
				start++
			}
			for start < end && nums[end] == nums[end-1] {
				end--
			}
			start++
			end--
		}
	}
	return sum
}

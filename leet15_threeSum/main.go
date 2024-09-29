package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func threeSum(nums []int) [][]int {

	sort.Ints(nums)
	var cur [][]int

	target := 0
	l := len(nums)
	for i := 0; i < l; i++ {
		res := twoSum(nums, i+1, target-nums[i])
		for _, r := range res {
			r = append(r, nums[i])
			cur = append(cur, r)
		}
		for i < l-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return cur
}

func twoSum(nums []int, start, target int) [][]int {
	end := len(nums) - 1

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

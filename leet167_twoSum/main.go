package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(twoSum(nums, 4))
}

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		}
		if numbers[left]+numbers[right] > target {
			right--
		}
		if numbers[left]+numbers[right] < target {
			left++
		}
	}
	return nil
}

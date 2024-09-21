package main

import "fmt"

func main() {
	nums := []int{1, 0, 2, 0, 3, 4}
	moveZeroes2(nums)
	fmt.Println(nums)
}

func moveZeros(nums []int) {
	p := removeEkement(nums, 0)
	for p < len(nums) {
		nums[p] = 0
		p++
	}
}

func removeEkement(nums []int, Val int) int {
	left := 0
	for _, v := range nums {
		if v != Val {
			nums[left] = v
			left++
		}
	}
	return left
}

func moveZeroes2(nums []int) {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
}

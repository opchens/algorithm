package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 1, 2, 2, 2, 3, 3}))
}

func removeDuplicates(nums []int) (ans int) {
	for left, right := 0, 0; left < len(nums); left = right {
		nums[ans] = nums[left]
		ans++
		for right < len(nums) && nums[right] == nums[left] {
			right++
		}
		if right-left > 1 {
			nums[ans] = nums[left]
			ans++
		}
	}
	return

}

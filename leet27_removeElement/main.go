package main

func removeElement(nums []int, val int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow

}

func removeElement2(nums []int, val int) int {
	left := 0
	for _, v := range nums {
		if nums[v] != val {
			nums[left] = nums[v]
			left++
		}
	}
	return left
}

package main

func majorityElement(nums []int) int {
	max := make([]int, 2)
	target := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		target[nums[i]]++
		if target[nums[i]] > max[1] {
			max[0] = nums[i]
			max[1] = target[nums[i]]
		}
	}
	return max[0]

}

func majorityElement2(nums []int) int {
	ans := nums[0]
	count := 0
	for _, num := range nums {
		if count == 0 {
			ans = num
			count++
		} else {
			if ans == num {
				count++
			} else {
				count--
			}
		}
	}
	return ans
}

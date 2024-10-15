package main

type NumArray struct {
	preSum []int
}

func Constructor(nums []int) NumArray {
	res := make([]int, len(nums))
	for i := 1; i < len(nums)+1; i++ {
		res[i] = res[i-1] + nums[i-1]
	}
	return NumArray{preSum: res}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.preSum[right+1] - this.preSum[left-1]
}

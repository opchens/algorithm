package main

type Difference struct {
	diff []int
}

func NewDifference(nums []int) *Difference {
	diff := make([]int, len(nums))
	diff[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		diff[i] = nums[i] + nums[i-1]
	}
	return &Difference{
		diff: diff,
	}
}

func (d *Difference) Increnment(update []int) {
	d.diff[update[0]] += update[2]
	if update[1]+1 < len(d.diff) {
		d.diff[update[1]+1] -= update[2]
	}
}

func (d *Difference) Result() []int {
	res := make([]int, len(d.diff))
	res[0] = d.diff[0]
	for i := 1; i < len(res); i++ {
		res[i] = res[i-1] + d.diff[i]
	}
	return res
}

func getModifiedArray(length int, updates [][]int) []int {
	nums := make([]int, length)
	for i := 0; i < length; i++ {
		nums[0] = 0
	}
	diff := NewDifference(nums)
	for i := 0; i < len(updates); i++ {
		diff.Increnment(updates[i])
	}
	return diff.Result()
}

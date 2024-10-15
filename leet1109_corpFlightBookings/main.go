package main

type Difference struct {
	diff []int
}

func NewDifference(nums []int) *Difference {
	diff := make([]int, len(nums))
	diff[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		diff[i] = nums[i] - nums[i-1]
	}
	return &Difference{
		diff: diff,
	}
}

func (d *Difference) Increment(i, j, value int) {
	d.diff[i] += value
	if j+1 < len(d.diff) {
		d.diff[j+1] -= value
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

func corpFlightBookings(bookings [][]int, n int) []int {
	nums := make([]int, n)
	for i := 0; i < len(nums); i++ {
		nums[i] = 0
	}
	diff := NewDifference(nums)
	for _, booking := range bookings {
		diff.Increment(booking[0], booking[1], booking[2])
	}
	return diff.Result()
}

package main

func main() {

}

func nthUglyNumber(n int) int {
	p2, p3, p5 := 1, 1, 1
	product2, product3, product5 := 1, 1, 1

	ugly := make([]int, n+1)
	p := 1
	for p <= n {
		m := min(product3, min(product5, product2))
		ugly[p] = m
		p++
		if m == product2 {
			product2 = 2 * ugly[p2]
			p2++
		}
		if m == product3 {
			product3 = 3 * ugly[p3]
			p3++
		}
		if m == product5 {
			product5 = 5 * ugly[p5]
			p5++
		}
	}
	return ugly[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

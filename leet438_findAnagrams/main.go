package main

func findAnagrams(s string, p string) []int {
	need, windows := make(map[byte]int), make(map[byte]int)
	left, right := 0, 0

	for i := range p {
		need[p[i]]++
	}

	valid := 0

	var res []int

	for right < len(s) {
		d := s[right]
		right++
		if _, ok := need[d]; ok {
			windows[d]++
			if windows[d] == need[d] {
				valid++
			}
		}

		for right-left >= len(p) {
			if valid == len(need) {
				res = append(res, left)
			}

			t := s[left]
			left++
			if _, ok := need[t]; ok {
				if windows[t] == need[t] {
					valid--
				}
				windows[t]--
			}

		}
	}
	return res
}

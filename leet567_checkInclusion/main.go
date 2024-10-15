package main

func main() {}

func checkInclusion(s1, s2 string) bool {
	need, windows := make(map[byte]int), make(map[byte]int)
	left, right := 0, 0

	for i := range s1 {
		need[s1[i]]++
	}

	valid := 0
	for right < len(s2) {
		s := s2[left]
		left++
		if _, ok := need[s]; ok {
			windows[s]++
			if windows[s] == need[s] {
				valid++
			}
		}

		for right-left >= len(s1) {
			if valid == len(need) {
				return true
			}

			v := s2[left]
			left++
			if _, ok := need[v]; ok {
				if need[v] == windows[v] {
					valid--
				}
				windows[v]--
			}
		}
	}
	return false
}

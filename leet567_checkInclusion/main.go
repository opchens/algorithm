package main

func main() {}

func checkInclusion(s1, s2 string) bool {
	need := make(map[byte]int, 0)
	window := make(map[byte]int, 0)
	for c := range s1 {
		need[s1[c]]++
	}

	left := 0
	valid := 0
	for right := 0; right < len(s2); right++ {
		c := s2[right]

	}

	return false
}

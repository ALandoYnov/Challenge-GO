package isnegative

func AlphaCount(s string) int {
	count := 0
	for _, n := range s {
		if (n >= 'a' && n <= 'z') || (n >= 'A' && n <= 'Z') {
			count++
		}
	}
	return count
}

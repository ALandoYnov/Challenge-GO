package isnegative

func NRune(s string, n int) rune {
	letter := []rune(s)
	if n > len(s) || n <= 0 {
		return 0
	} else {
	}
	return letter[n-1]
}

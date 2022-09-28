package isnegative

func IsPrintable(s string) bool {
	phrase := []rune(s)
	counter := 0
	for i := 0; i < len(s); i++ {
		if phrase[i] >= 32 && phrase[i] <= 126 {
			counter++
			if len(s) == counter {
				return true
			}
		}
	}
	return false
}

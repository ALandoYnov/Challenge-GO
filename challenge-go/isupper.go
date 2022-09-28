package isnegative

func IsUpper(s string) bool {
	phrase := []rune(s)
	counter := 0
	for i := 0; i < len(s); i++ {
		if phrase[i] >= 65 && phrase[i] <= 90 {
			counter++
			if len(s) == counter {
				return true
			}
		}
	}
	return false
}

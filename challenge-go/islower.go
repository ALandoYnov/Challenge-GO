package isnegative

func IsLower(s string) bool {
	phrase := []rune(s)
	counter := 0
	for i := 0; i < len(s); i++ {
		if phrase[i] >= 97 && phrase[i] <= 122 {
			counter++
			if len(s) == counter {
				return true
			}
		}
	}
	return false
}

package isnegative

func ToUpper(s string) string {
	b := make([]byte, len(s))
	for i := range b {
		c := s[i]
		if c >= 'a' && c <= 'z' {
			c -= 'a' - 'A'
		}
		b[i] = c
	}
	return string(b)
}

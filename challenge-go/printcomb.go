package isnegative

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	for a := '0'; a <= '9'; a++ {
		for b := '0'; b <= '9'; b++ {
			for c := '0'; c <= '9'; c++ {
				if a < b {
					if b < c {
						z01.PrintRune(a)
						z01.PrintRune(b)
						z01.PrintRune(c)
						if a == '7' && b == '8' && c == '9' {
							z01.PrintRune(10)
						} else {
							z01.PrintRune(44)
							z01.PrintRune(32)
						}
					}
				}
			}
		}
	}
}

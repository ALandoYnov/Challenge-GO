package isnegative

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	for a := 0; a <= 99; a++ {
		for b := a + 1; b <= 99; b++ {
			da := a / 10
			ua := a % 10
			db := b / 10
			ub := b % 10
			// a := 54
			// da -> 5
			z01.PrintRune(rune('0' + da))
			z01.PrintRune(rune('0' + ua))
			z01.PrintRune(32)
			z01.PrintRune(rune('0' + db))
			z01.PrintRune(rune('0' + ub))
			if a == 98 && b == 99 {
				z01.PrintRune(10)
			} else {
				z01.PrintRune(44)
				z01.PrintRune(32)
			}
		}
	}
}

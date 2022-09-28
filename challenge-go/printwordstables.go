package isnegative

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	var resultat string
	for i := range a {
		resultat += a[i]
		if i < len(a) {
			resultat += "\n"
		}
	}
	for i := range resultat {
		z01.PrintRune(rune(resultat[i]))
	}
}

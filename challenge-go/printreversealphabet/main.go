package main

import "github.com/01-edu/z01"

func main() {
	var alphabet rune = 'z'
	for alphabet >= 'a' {
		z01.PrintRune(alphabet)
		alphabet--
	}
	z01.PrintRune('\n')
}

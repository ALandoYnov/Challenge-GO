package main

import "github.com/01-edu/z01"

func main() {
	var number rune = '0'
	for number <= '9' {
		z01.PrintRune(number)
		number++
	}
	z01.PrintRune('\n')
}

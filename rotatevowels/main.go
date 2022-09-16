package main

import (
	"os"

	"github.com/01-edu/z01"
)

func check(x rune) bool {
	if x == 'a' || x == 'A' || x == 'e' || x == 'E' || x == 'o' || x == 'O' || x == 'u' || x == 'U' || x == 'i' || x == 'I' {
		return true
	}
	return false
}
func main() {
	arg := os.Args[1:]
	table := []rune{}
	runeStr := ""
	taille := 0
	IsFalse := true
	for _, c := range arg {
		for _, j := range c {
			if check(j) {
				table = append(table, j)
				taille++
			}
		}
		if IsFalse {
			runeStr = c
			IsFalse = false
			continue
		}
		runeStr = runeStr + " " + c
	}
	cur := 0
	for _, c := range runeStr {
		if check(c) {
			z01.PrintRune(table[taille-cur-1])
			cur++
		} else {
			z01.PrintRune(c)
		}
	}
	z01.PrintRune('\n')
}

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	table := os.Args[1:]
	taille := 0
	for i := range table {
		taille = i
	}
	if taille >= 1 {
		if table[0] == "--upper" {
			for i := 1; i <= taille; i++ {
				nb := 0
				for _, w := range table[i] {
					nb = nb*10 + int(w-'0')
				}
				if nb >= 1 && nb <= 26 {
					z01.PrintRune('A' + rune(nb-1))
				} else {
					z01.PrintRune(' ')
				}
			}
		} else {
			for i := 0; i <= taille; i++ {
				numb := 0
				for _, w := range table[i] {
					numb = numb*10 + int(w-'0')
				}
				if numb >= 1 && numb <= 26 {
					z01.PrintRune('a' + rune(numb-1))
				} else {
					z01.PrintRune(' ')
				}
			}
		}
	}
	if taille <= 0 {
		z01.PrintRune('\n')
	}
}

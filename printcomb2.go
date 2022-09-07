package student

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	for a := 0; a <= 9; a++ {
		for b := 0; b <= 9; b++ {
			for i := a; i <= 9; i++ {
				for j := 0; j <= 9; j++ {
					if (a == i && b == j) || (i <= a && j <= b) {
						continue
					}
					z01.PrintRune(rune(a) + 48)
					z01.PrintRune(rune(b) + 48)
					z01.PrintRune(' ')
					z01.PrintRune(rune(i) + 48)
					z01.PrintRune(rune(j) + 48)
					if !(a == 9 && b == 8 && i == 9 && j == 9) {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}

package student

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	for a := 0; a < 10; a++ {
		for b := 0; b < 10; b++ {
			for i := 0; i < 10; i++ {
				for j := 0; j < 10; j++ {
					if a+b<i+j {
						z01.PrintRune(rune(a) + 48)
						z01.PrintRune(rune(b) + 48)
						z01.PrintRune(' ')
						z01.PrintRune(rune(i) + 48)
						z01.PrintRune(rune(j) + 48)
						if rune(a+48) != '9' || rune(b+48) != '9' || rune(i+48) != '9' || rune(j+48) != '9' {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						} else {
							z01.PrintRune('\n')
						}
					}
				}
			}
		}
	}
}

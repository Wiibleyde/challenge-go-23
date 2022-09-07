package student

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	negative := n < 0
	if negative && (n*-1) != 0 {
		n *= -1
	}
	if negative {
		z01.PrintRune('-')
	}
	rec(n)
}

func rec(n int) {
	for i := int(0); i <= 9; i++ {
		if n <= 9 && n >= -9 {
			z01.PrintRune(rune(n + 48))
			return
		}
		if (n-i)%10 == 0 {
			rec((n - i) / 10)
			z01.PrintRune(rune(i + 48))
			return
		}
	}
}

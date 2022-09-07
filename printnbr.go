package student

import (
	"github.com/01-edu/z01"
)

func reverse_int(n int) int {
	new_int := 0
	negative := false
	if n < 0 {
		negative = true
		n = n * -1
	}
    for n > 0 {
        remainder := n % 10
        new_int *= 10
        new_int += remainder
        n /= 10
	}
	if negative {
		return new_int * -1
	} else {
		return new_int
	}
}

func PrintNbr(n int) {
	nb := reverse_int(n)
	if nb < 0 {
		z01.PrintRune('-')
		nb = nb * -1
	}
	if nb == 0 {
		z01.PrintRune('0')
	}
	for nb > 0 {
		z01.PrintRune(rune(nb % 10 + '0'))
		nb = nb / 10
	}
	z01.PrintRune('\n')
}

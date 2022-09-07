package main

import (
	"github.com/01-edu/z01"
)

func main() {
	PrintNbr(-9223372036854775808)
	z01.PrintRune('\n')
}

// func reverse_int(n int) int {
// 	new_int := 0
// 	negative := false
// 	if n < 0 {
// 		negative = true
// 		n = n * -1
// 	}
// 	for n > 0 {
// 		remainder := n % 10
// 		new_int *= 10
// 		new_int += remainder
// 		n /= 10
// 	}
// 	if negative {
// 		return new_int * -1
// 	} else {
// 		return new_int
// 	}
// }

func PrintNbr(n int) {
	negative := n < 0
	if negative {
		z01.PrintRune('-')
	}

	rec(n)
}

func rec(n int) {
	for i := int(0); i < 10; i++ {
		if n <= 9 && n >= -9 {
			z01.PrintRune(rune(n + 48))
			return
		}
		if (n-i)%10 == 0 {
			rec((n-i)/10)
			z01.PrintRune(rune(i + 48))
			return
		}
	}
}

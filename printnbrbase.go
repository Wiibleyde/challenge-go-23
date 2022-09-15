package student

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(n int, base string) {
	value := ""
	lenStr := 0
	for _, counter := range base {
		_ = counter
		lenStr++
	}
	maxLen := lenStr
	if n < 0 {
		value = "-"
		maxLen *= -1
	}
	if lenStr > 1 {
		for n/maxLen >= lenStr {
			maxLen *= lenStr
		}
		for maxLen != 0 {
			value = value + string(base[n/maxLen])
			n = n - n/maxLen*maxLen
			maxLen /= lenStr
		}
		x := map[rune]bool{}
		for _, counter := range base {
			if counter == '+' || counter == '-' {
				value = "NV"
				break
			}
			if !x[counter] {
				x[counter] = true
			} else {
				value = "NV"
				break
			}
		}
	} else {
		value = "NV"
	}
	for _, counter := range value {
		z01.PrintRune(counter)
	}
}

package student

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, ele := range a {
		for i := 0; i < StrLen(ele); i++ {
			z01.PrintRune(rune(ele[i]))
		}
		z01.PrintRune('\n')
	}
}

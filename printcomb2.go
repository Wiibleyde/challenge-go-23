package main

import (
	"github.com/01-edu/z01"
)

func main() {
	PrintComb2()
}

func PrintComb2 () {
	for i := 0; i < 10; i++{
		for j := 0; j < 10; j++{
			z01.PrintRune(rune(i)+48)
			z01.PrintRune(rune(j)+48)
			if rune(i+48) != '9' || rune(j+48) != '9' {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			} else {
				z01.PrintRune('\n')
			}
		}
	}
}

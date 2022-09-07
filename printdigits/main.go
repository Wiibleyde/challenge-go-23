package main

import "github.com/01-edu/z01"

func main() {
	numbers := "0123456789"
	for _, char := range numbers {
		z01.PrintRune(rune(char))
	}
	z01.PrintRune('\n')
}

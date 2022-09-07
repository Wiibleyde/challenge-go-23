package main

import "github.com/01-edu/z01"

func main() {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	for _, char := range alphabet {
		z01.PrintRune(rune(char))
	}
	z01.PrintRune('\n')
}

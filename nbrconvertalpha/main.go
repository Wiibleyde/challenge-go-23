package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	for _, ele := range args {
		for _, ele2 := range ele {
			z01.PrintRune(rune(byte(ele2)))
		}
		z01.PrintRune('\n')
	}
}

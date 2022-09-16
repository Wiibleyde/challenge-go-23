package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	for _, ele := range args[1:] {
		for _, ele2 := range ele {
			z01.PrintRune(ele2)
		}
		z01.PrintRune('\n')
	}
}

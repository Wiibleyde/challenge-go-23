package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		for j := i; j > 0 && args[j-1] > args[j]; j-- {
			args[j], args[j-1] = args[j-1], args[j]
		}
	}
	for _, ele := range args {
		for _, ele2 := range ele {
			z01.PrintRune(ele2)
		}
		z01.PrintRune('\n')
	}
}

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	args := ""
	vowels := []string{}
	result := ""
	changeEffectued := false
	if len(arguments) < 1 {
		z01.PrintRune('\n')
	}
	for i := 0; i < len(arguments); i++ {
		args = args + arguments[i] + " "
	}
	for i := 0; i < len(args); i++ {
		if args[i] == 'a' || args[i] == 'e' || args[i] == 'i' || args[i] == 'o' || args[i] == 'u' || args[i] == 'A' || args[i] == 'E' || args[i] == 'I' || args[i] == 'O' || args[i] == 'U' {
			vowels = append(vowels, string(args[i]))
			changeEffectued = true
		}
	}
	for i := 0; i < len(vowels); i++ {
		for j := i + 1; j < len(vowels); j++ {
			vowels[i], vowels[j] = vowels[j], vowels[i]
			changeEffectued = true
		}
	}
	j := 0
	for i := 0; i < len(args); i++ {
		if args[i] == 'a' || args[i] == 'e' || args[i] == 'i' || args[i] == 'o' || args[i] == 'u' || args[i] == 'A' || args[i] == 'E' || args[i] == 'I' || args[i] == 'O' || args[i] == 'U' {
			result = result + vowels[j]
			changeEffectued = true
			j++

		} else {
			result = result + string(args[i])
			changeEffectued = true
		}
	}
	for i := 0; i < len(result); i++ {
		z01.PrintRune(rune(result[i]))
		changeEffectued = true
	}
	if changeEffectued {
		z01.PrintRune('\n')
	}
}

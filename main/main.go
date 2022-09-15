package main

import (
	"student"

	"github.com/01-edu/z01"
)

func main() {
	z01.PrintRune(student.NRune("Hello!", 3))
	z01.PrintRune(student.NRune("Salut!", 2))
	z01.PrintRune(student.NRune("Bye!", -1))
	z01.PrintRune(student.NRune("Bye!", 5))
	z01.PrintRune(student.NRune("Ola!", 4))
	z01.PrintRune('\n')
}

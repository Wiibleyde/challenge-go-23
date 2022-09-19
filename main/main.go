package main

import (
	"student"
)

func main() {
	a := student.SplitWhiteSpaces("Hello how are you?")
	student.PrintWordsTables(a)
}

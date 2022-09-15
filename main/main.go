package main

import (
	"fmt"
	"student"
)

func main() {
	fmt.Println(student.IsAlpha("Hello! How are you?"))
	fmt.Println(student.IsAlpha("HelloHowareyou"))
	fmt.Println(student.IsAlpha("What's this 4?"))
	fmt.Println(student.IsAlpha("Whatsthis4"))
}

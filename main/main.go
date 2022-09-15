package main

import (
	"fmt"
	"student"
)

func main() {
	fmt.Println(student.IsPrintable("s\\d Gs4kT[z/`"))
	fmt.Println(student.IsPrintable("Hello\n"))
}

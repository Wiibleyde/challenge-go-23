package main

import (
	"fmt"
	"student"
)

func main() {
	fmt.Println(student.IsPrintable("Hello"))
	fmt.Println(student.IsPrintable("Hello\n"))
}

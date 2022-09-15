package main

import (
	"fmt"
	"student"
)

func main() {
	fmt.Println(student.IsNumeric("010203"))
	fmt.Println(student.IsNumeric("01,02,03"))
}

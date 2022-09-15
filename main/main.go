package main

import (
	"student"
	"fmt"
)

func main() {
	fmt.Println(student.Compare("Hello!", "Hello!"))
	fmt.Println(student.Compare("Salut!", "lut!"))
	fmt.Println(student.Compare("Ola!", "Ol"))
}

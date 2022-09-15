package main

import (
	"student"

	"fmt"
	
	// "github.com/01-edu/z01"
)

func main() {
	fmt.Println(student.Compare("Hello!", "Hello!"))
	fmt.Println(student.Compare("Salut!", "lut!"))
	fmt.Println(student.Compare("Ola!", "Ol"))
}

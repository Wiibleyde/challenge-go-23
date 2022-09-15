package main

import (
	"fmt"
	"student"
)

func main() {
	fmt.Println(student.Index("Hello!", "l"))
	fmt.Println(student.Index("Salut!", "alu"))
	fmt.Println(student.Index("Ola!", "hOl"))
}

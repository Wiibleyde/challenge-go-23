package main

import (
	"fmt"
	"student"
)

func main() {
	elems := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(student.BasicJoin(elems))
}

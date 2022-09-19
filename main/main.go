package main

import (
	"fmt"
	"student"
)

func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(student.ConcatParams(test))
}

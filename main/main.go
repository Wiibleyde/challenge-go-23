package main

import (
	"fmt"
	"student"
)

func main() {
	s := "Hello 78 World!    4455 /"
	nb := student.AlphaCount(s)
	fmt.Println(nb)
}

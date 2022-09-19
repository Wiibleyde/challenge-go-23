package main

import (
	"fmt"
	"student"
)

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", student.Split(s, "HA"))
}

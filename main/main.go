package main

import "student"

func main () {
	var a int
	var b int
	a = 1
	b = 0
	student.Swap(&a, &b)
	println(a)
	println(b)	
}
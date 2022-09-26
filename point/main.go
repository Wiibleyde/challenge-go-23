package main

import "github.com/01-edu/z01"

type point struct {
	x, y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func main() {
	points := &point{}

	setPoint(points)

	IntY := Itoa(points.x)
	IntX := Itoa(points.y)

	printStr("x = " + string(IntY) + ", y = " + string(IntX))
}

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var res string
	for n != 0 {
		res = string(n%10+48) + res
		n /= 10
	}
	return res
}

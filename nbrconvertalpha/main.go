package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if stringArrLen(os.Args) <= 1 {
		z01.PrintRune('\n')
		return
	}
	for _, s := range os.Args[1:] {
		nm := atoi(s)
		if nm <= 26 && nm > 0 {
			if os.Args[1] == "--upper" {
				z01.PrintRune(rune(nm + 96 - 32))
			} else {
				z01.PrintRune(rune(nm + 96))
			}
		} 
	}
	z01.PrintRune('\n')
}

func stringArrLen(strs []string) int {
	count := 0
	for range strs {
		count++
	}

	return count
}

func strLen(str string) int {
	i := 0
	for range str {
		i++
	}
	return i
}

func atoi(s string) int {
	if strLen(s) == 0 {
		return 0
	}
	s0 := s
	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
		if strLen(s) < 1 {
			return 0
		}
	}
	nm := 0
	for _, ch := range s {
		if !containsIn0to9(ch) {
			return 0
		}
		nm = nm*10 + charToInt(ch)
	}
	if s0[0] == '-' {
		nm *= -1
	}
	return nm
}

func charToInt(ch rune) int {
	count := 0
	if ch < 48 && ch > 58 {
		return 0
	}
	for i := '1'; i <= ch; i++ {
		count++
	}
	return count
}

func containsIn0to9(ch rune) bool {
	for i := '0'; i <= '9'; i++ {
		if ch == i {
			return true
		}
	}
	return false
}

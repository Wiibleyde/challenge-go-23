package main

import (
	"os"
)

func Atoi(s string) int {
	if StrLen(s) == 0 {
		return 0
	}
	s0 := s
	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
		if StrLen(s) < 1 {
			return 0
		}
	}
	nb := 0
	for _, ch := range s {
		if !containsIn0to9(ch) {
			return 0
		}
		nb = nb*10 + charToInt(ch)
	}
	if s0[0] == '-' {
		nb *= -1
	}
	return nb
}

func StrLen(s string) int {
	letter := 0
	for _, ele := range s {
		letter++
		_ = ele
	}
	return letter
}

func charToInt(char rune) int {
	count := 0
	if char < 48 && char > 58 {
		return 0
	}

	for i := '1'; i <= char; i++ {
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

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	negative := false
	if n < 0 {
		negative = true
		n = n * -1
	}
	var res string
	for n != 0 {
		res = string(n%10+48) + res
		n /= 10
	}
	if negative {
		return "-" + res
	}
	return res
}

func main() {
	args := os.Args[1:]
	a1 := 0
	a2 := 0
	if len(args) != 3 {
		return
	}
	if !(IsNumeric(args[0]) && IsNumeric(args[2])) {
		return
	}
	a1 = Atoi(args[0])
	a2 = Atoi(args[2])
	if a1 >= 9223372036800000000 || a2 >= 9223372036800000000 {
		return
	}
	switch args[1] {
	case "*":
		Display(a1 * a2)
	case "/":
		if a2 == 0 {
			os.Stdout.WriteString("No division by 0" + "\n")
			return
		}
		Display(a1 / a2)
	case "%":
		if a2 == 0 {
			os.Stdout.WriteString("No modulo by 0" + "\n")
			return
		}
		Display(a1 % a2)
	case "+":
		Display(a1 + a2)
	case "-":
		Display(a1 - a2)
	}
}

func Display(i int) {
	if i >= 9223372036854775806 {
		return
	}
	str := Itoa(i)
	if str == "-" || str == "" || str == " " {
		return
	}
	os.Stdout.WriteString(str + "\n")
}

func IsNumeric(s string) bool {
	for _, c := range s {
		if !(c >= 48 && c <= 57) && c != '-' {
			return false
		}
	}
	return true
}

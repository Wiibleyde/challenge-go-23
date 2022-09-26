package main

import (
	"os"
)

// func main() {
// 	args := os.Args[1:]
// 	if len(args) == 3 {
// 		// check if the first argument is a number and the second to positive or negative
// 		if containsIn0to9(rune(args[0][0])) && (args[1] == "+" || args[1] == "-" || args[1] == "*" || args[1] == "/" || args[1] == "%") && containsIn0to9(rune(args[2][0])) {
// 			if args[1] == "+" {
// 				os.Stdout.WriteString(Itoa(sum(Atoi(args[0]), Atoi(args[2]))))
// 				os.Stdout.WriteString("\n")
// 			} else if args[1] == "-" {
// 				os.Stdout.WriteString(Itoa(substract(Atoi(args[0]), Atoi(args[2]))))
// 				os.Stdout.WriteString("\n")
// 			} else if args[1] == "*" {
// 				os.Stdout.WriteString(Itoa(multiply(Atoi(args[0]), Atoi(args[2]))))
// 				os.Stdout.WriteString("\n")
// 			} else if args[1] == "/" {
// 				if Atoi(args[2]) == 0 {
// 					os.Stdout.WriteString("No division by 0")
// 					os.Stdout.WriteString("\n")
// 				} else {
// 					os.Stdout.WriteString(Itoa(divide(Atoi(args[0]), Atoi(args[2]))))
// 					os.Stdout.WriteString("\n")
// 				}
// 			} else if args[1] == "%" {
// 				if Atoi(args[2]) == 0 {
// 					os.Stdout.WriteString("No modulo by 0")
// 					os.Stdout.WriteString("\n")
// 				} else {
// 					os.Stdout.WriteString(modulo(Atoi(args[0]), Atoi(args[2])))
// 					os.Stdout.WriteString("\n")
// 				}
// 			}
// 		}
// 	}
// }

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

// func sum(a, b int) int {
// 	return a + b
// }

// func substract(a, b int) int {
// 	return a - b
// }

// func multiply(a, b int) int {
// 	return a * b
// }

// func divide(a, b int) int {
// 	return a / b
// }

// func modulo(a, b int) string {
// 	if a == 0 || b == 0 {
// 		os.Stdout.WriteString("No modulo by 0")
// 	}
// 	return Itoa(a % b)
// }

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

func isSigne(str string, arr []string) bool {
	for _, v := range arr {
		if str == v {
			return true
		}
	}
	return false
}

func main() {
	signe := []string{"+", "*", "-", "/", "%"}
	args := os.Args[1:]
	if len(args) == 3 {
		if isSigne(args[1], signe) {
			if containsIn0to9(rune(args[0][0])) && (args[1] == "+" || args[1] == "-" || args[1] == "*" || args[1] == "/" || args[1] == "%") && containsIn0to9(rune(args[2][0])) {
				nb1 := Atoi(args[0])
				nb2 := Atoi(args[2])
				switch args[1] {
				case "+":
					os.Stdout.WriteString(Itoa(nb1 + nb2))
				case "-":
					os.Stdout.WriteString(Itoa(nb1 - nb2))
				case "/":
					if nb2 == 0 {
						os.Stdout.WriteString("No division by 0")
					} else {
						os.Stdout.WriteString(Itoa(nb1 / nb2))
					}
				case "%":
					if nb2 == 0 {
						os.Stdout.WriteString("No Modulo by 0")
					} else {
						os.Stdout.WriteString(Itoa(nb1 % nb2))
					}
				case "*":
					os.Stdout.WriteString(Itoa(nb1 * nb2))
				}
			} else {
				os.Stdout.WriteString("1")
			}
		} else {
			os.Stdout.WriteString("0")
		}
		os.Stdout.WriteString("\n")
	}
}

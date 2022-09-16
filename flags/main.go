package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func order(s string) {
	var a [5000]int
	for _, c := range s {
		a[int(c)]++
	}
	for i, c := range a {
		for c > 0 {
			z01.PrintRune(rune(i))
			c--
		}
	}
	z01.PrintRune('\n')
}

func insert(str1 string, str2 string) string {
	return str1 + str2
}

func help() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("	 This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("	 This flag will behave like a boolean, if it is called it will order the argument.")
}

func main() {
	args := os.Args[1:]
	ln := -1
	for i := range args {
		ln = i
	}
	if ln != -1 {
		if args[0] == "-h" || args[0] == "--help" {
			help()
		} else if args[0] == "--order" || args[0] == "-o" {
			order(args[1])
		} else if args[0][0:3] == "--i" || args[0][0:2] == "-i" {
			if ln < 2 {
				if args[0][0:3] == "--i" {
					fmt.Println(insert(args[1], args[0][9:]))
				} else {
					fmt.Println(insert(args[1], args[0][3:]))
				}
			} else {
				ele := ""
				for i := 1; i <= ln; i++ {
					if args[i] != "-o" && args[i] != "--order" {
						ele = insert(ele, args[i])
					}
				}
				if args[0][0:3] == "--i" {
					ele = insert(ele, args[0][9:])
				} else {
					ele = insert(ele, args[0][3:])
				}
				order(ele)
			}
		} else {
			fmt.Println(args[0])
		}
	} else {
		help()
	}
}

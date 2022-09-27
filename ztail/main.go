package main

import (
	"fmt"
	"os"
)

func numberOfBytes(args []string) (int, []string) {
	n := len(args)
	nbytes := 0
	var files []string
	for i, v := range args {
		if v == "-c" {
			if i+1 < n {
				nbytes = Atoi(args[i+1])
				files = args[i+2:]
				break
			}
		}
	}
	return nbytes, files
}

func fileSize(fi *os.File) int64 {
	fil, err := fi.Stat()
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return fil.Size()
}

func main() {
	n := len(os.Args)
	if n < 4 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}
	nbytes, files := numberOfBytes(os.Args[1:])
	printName := len(files) > 1
	for j, f := range files {
		fi, err := os.Open(f)
		if err != nil {
			fmt.Printf("tail: cannot open '%s' for reading: No such file or directory\n", f)
			os.Exit(1)
		}
		if printName {
			fmt.Printf("==> %s <==\n", f)
		}
		read := make([]byte, int(nbytes))
		_, er := fi.ReadAt(read, fileSize(fi)-int64(nbytes))
		if er != nil {
			fmt.Println(er.Error())
		}
		for _, c := range read {
			fmt.Print(c)
		}
		if j < len(files)-1 {
			fmt.Print("\n")
		}
		fi.Close()
	}
}

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

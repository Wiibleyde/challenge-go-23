package main

import (
	"fmt"
	"os"
)

func numberOfBytes(args []string) (int, []string) {
	nbytes := 10
	files := []string{}
	for _, s := range args {
		if containsIn0to9(rune(s[0])) {
			nbytes = 0
			for _, r := range s {
				if containsIn0to9(r) {
					nbytes = nbytes*10 + int(r-'0')
				} else {
					break
				}
			}
		} else {
			files = append(files, s)
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
	args := os.Args[1:]
	nbytes, files := numberOfBytes(args)
	if len(files) == 0 {
		files = append(files, "")
	}
	for _, s := range files {
		fi, err := os.Open(s)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer fi.Close()
		size := fileSize(fi)
		if nbytes > int(size) {
			nbytes = int(size)
		}
		buf := make([]byte, nbytes)
		fi.ReadAt(buf, size-int64(nbytes))
		fmt.Print(string(buf))
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

package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		printStr("ERROR: open asd: no such file or directory")
	} else {
		for _, s := range os.Args[1:] {
			file, err := os.Open(s)
			if err != nil {
				printStr(err.Error())
				break
			} else {
				data, err := ioutil.ReadAll(file)
				if err != nil {
					printStr(err.Error())
					break
				} else {
					printStr(string(data))
					z01.PrintRune('\n')
				}
			}
		}
	}
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
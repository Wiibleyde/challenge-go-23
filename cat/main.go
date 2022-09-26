package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if 0 == len(args) {
		printStr("File name missing")
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
					printStr(convertBytesToString(data))
				}
			}
		}
	}
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func convertBytesToString(data []byte) string {
	return string(data[:])
}

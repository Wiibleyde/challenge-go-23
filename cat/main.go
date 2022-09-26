package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	errorCount := 0
	if len(args) != 0 {
		for _, s := range os.Args[1:] {
			data, err := ioutil.ReadFile(s)
			if err != nil {
				printStr("ERROR: " + err.Error())
				break
			} else {
				printStr(convertBytesToString(data))
			}
		}
		if errorCount > 0 {
			os.Exit(1)
		}
	}
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func convertBytesToString(data []byte) string {
	return string(data[:])
}

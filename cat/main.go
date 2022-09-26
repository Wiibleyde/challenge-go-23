package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 0 {
		for _, s := range os.Args[1:] {
			data, err := ioutil.ReadFile(s)
			if err != nil {
				printStr("ERROR: " + err.Error() + "\n")
				os.Exit(1)
				break
			}
			printStr(string(data))
		}
	} else {
		io.Copy(os.Stdout, os.Stdin)
	}
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

// func convertBytesToString(data []byte) string {
// 	return string(data[:])
// }

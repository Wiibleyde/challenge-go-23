// package main

// import (
// 	"fmt"
// 	"os"
// )

// func numberOfBytes(args []string) (int, []string) {
// 	nbytes := 10
// 	files := []string{}
// 	for _, s := range args {
// 		if containsIn0to9(rune(s[0])) {
// 			nbytes = 0
// 			for _, r := range s {
// 				if containsIn0to9(r) {
// 					nbytes = nbytes*10 + int(r-'0')
// 				} else {
// 					break
// 				}
// 			}
// 		} else {
// 			files = append(files, s)
// 		}
// 	}
// 	return nbytes, files
// }

// func fileSize(fi *os.File) int64 {
// 	fil, err := fi.Stat()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return 0
// 	}
// 	return fil.Size()
// }

// func main() {
// 	n := len(os.Args)
// 	if n < 4 {
// 		fmt.Println("Not enough arguments")
// 		os.Exit(1)
// 	}
// 	nbytes, files := numberOfBytes(os.Args[1:])
// 	printName := len(files) > 1
// 	for j, f := range files {
// 		fi, err := os.Open(f)
// 		if err != nil {
// 			fmt.Printf("open %s: no such file or directory\n", f)
// 			fmt.Printf("\n")
// 			defer os.Exit(1)
// 			continue
// 		}
// 		if printName {
// 			if j > 0 || j < len(files)-1 {
// 				fmt.Printf("==> %s <==\n", f)
// 			} else {
// 				fmt.Printf("==> %s <==\n", f)
// 			}
// 		}
// 		read := make([]byte, int(nbytes))
// 		_, err = fi.Read(read)
// 		if err != nil {
// 			fmt.Println(err.Error())
// 		}
// 		fmt.Printf("%s", read)
// 		fmt.Printf("\n")
// 	}
// }

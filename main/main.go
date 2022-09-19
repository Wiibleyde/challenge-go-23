package main

import (
	"fmt"
	"student"
)

func main() {
	result := student.ConvertBase("1010110", "01", "0123456789")
	fmt.Println(result)
}

package main

func IsNegative(nb int) {
	if nb < 0 {
		print(true)
		print("\n")
	} else {
		print(false)
		print("\n")
	}
}

func main() {
	IsNegative(1)
	IsNegative(0)
	IsNegative(-1)
}

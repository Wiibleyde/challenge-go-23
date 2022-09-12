package student

import "github.com/01-edu/z01"

const N = 8

var position = [N]int{}

func Calc(queenpos, rowposition int) bool {
	for i := 0; i < queenpos; i++ {
		other_row_pos := position[i]
		if other_row_pos == rowposition || other_row_pos == rowposition-(queenpos-i) || other_row_pos == rowposition+(queenpos-i) {
			return false
		}
	}
	return true
}

func Resolve(k int) {
	if k == N {
		for i := 0; i < N; i++ {
			printer(position[i] + 1)
		}
		z01.PrintRune('\n')
	} else {
		for i := 0; i < N; i++ {
			if Calc(k, i) {
				position[k] = i
				Resolve(k + 1)
			}
		}
	}
}

func EightQueens() {
	Resolve(0)
}

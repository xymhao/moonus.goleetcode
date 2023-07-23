package _struct

var result [8]int

func init() {
	result = [8]int{-1, -1, -1, -1, -1, -1, -1, -1}
}

func eightQueens(row int) {
	if row == 8 {
		printQueens()
		return
	}
	for column := 0; column < 8; column++ {
		if isOk(row, column) {
			result[row] = column
			eightQueens(row + 1)
		}
	}
}

func isOk(row, column int) bool {
	leftUp, rightUp := column-1, column+1
	for i := row - 1; i >= 0; i-- {
		if result[i] == column {
			return false
		}
		if leftUp >= 0 && result[i] == leftUp {
			return false
		}
		if rightUp < 8 && result[i] == rightUp {
			return false
		}
		leftUp--
		rightUp++
	}
	return true
}

func printQueens() {
	for row := 0; row < 8; row++ {
		for column := 0; column < 8; column++ {
			if result[row] == column {
				print("Q ")
			} else {
				print("* ")
			}
		}
		println()
	}
	println()
}

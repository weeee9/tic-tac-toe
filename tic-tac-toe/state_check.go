package tictactoe

func checkWin(s symbol, blocks [3][3]*Block) bool {
	if s == symbolNon {
		return false
	}

	if checkRow(s, blocks) {
		return true
	}
	if checkCol(s, blocks) {
		return true
	}
	return checkDia(s, blocks)
}

func checkRow(s symbol, blocks [3][3]*Block) bool {
	for _, row := range blocks {
		if row[0].Value() == s && row[1].Value() == s && row[2].Value() == s {
			return true
		}
	}
	return false
}

func checkCol(s symbol, blocks [3][3]*Block) bool {
	for i := 0; i < 3; i++ {
		if blocks[0][i].Value() == s && blocks[1][i].Value() == s && blocks[2][i].Value() == s {
			return true
		}
	}
	return false
}

func checkDia(s symbol, blocks [3][3]*Block) bool {
	// check main diag
	if blocks[0][0].Value() == s && blocks[1][1].Value() == s && blocks[2][2].Value() == s {
		return true
	}

	// check anit-diag
	if blocks[0][2].Value() == s && blocks[1][1].Value() == s && blocks[2][0].Value() == s {
		return true
	}
	return false
}

package leetcode

func IsValidSudoku(board [][]byte) bool {
	mp := make(map[byte][]int, 0)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				mp[board[i][j]] = append(mp[board[i][j]], i*9+j)
			}
		}
	}
	for _, pos := range mp {
		for i := 0; i < len(pos); i++ {
			for j := i + 1; j < len(pos); j++ {
				if (pos[j]-pos[i])%9 == 0 {
					return false
				}
				if pos[j]/9 == pos[i]/9 {
					return false
				}
				if pos[j]%9/3 == pos[i]%9/3 && pos[j]/9/3 == pos[i]/9/3 {
					return false
				}
			}
		}
	}
	return true
}

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' && !checkValid(board[i][j], []int{i, j}, board) {
				return false
			}
		}
	}
	return true
}

func checkValid(num byte, pos []int, board [][]byte) bool {
	row, column := pos[0], pos[1]
	a, b := row/3*3, column/3*3
	for i := 0; i < 9; i++ {
		if board[row][i] == num && i != column {
			return false
		}
		if board[i][column] == num && i != row {
			return false
		}
		if board[a+i/3][b+i%3] == num && a+i/3 != row && b+i%3 != column {
			return false
		}
	}
	return true
}

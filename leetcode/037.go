package leetcode

//DFS
func SolveSudoku(board [][]byte) {
	ans := []byte{49, 50, 51, 52, 53, 54, 55, 56, 57}
	poses := make([][]int, 0)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				poses = append(poses, []int{i, j})
			}
		}
	}
	for _, b := range ans {
		fill(b, poses, board, 0)
	}
	return
}

func fill(num byte, poses [][]int, board [][]byte, index int) bool {
	if index == len(poses) {
		return true
	}
	ans := []byte{49, 50, 51, 52, 53, 54, 55, 56, 57}
	pos := poses[index]
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
	board[row][column] = num
	for _, b := range ans {
		if fill(b, poses, board, index+1) {
			return true
		}
	}
	board[row][column] = '.'
	return false
}

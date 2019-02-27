package main

import "fmt"

var num = 0

func main() {
	//fmt.Println(test("abacab"))
	fmt.Println(search([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "ABCB"))

}

/**
查找单词
*/
func search(board [][]byte, word string) bool {
	w := []byte(word)
	width, length := len(board[0]), len(board)
	flag := make([][]int, length)
	for i := 0; i < len(flag); i++ {
		flag[i] = make([]int, width)
	}
	if width*length < len(w) {
		return false
	}
	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			if dfs(w, board, i, j, &flag, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(w []byte, board [][]byte, i int, j int, flag *[][]int, index int) bool {
	width, length := len(board[0]), len(board)
	if i >= length || j >= width || i < 0 || j < 0 || board[i][j] != w[index] {
		return false
	} else {
		if (*flag)[i][j] == 0 {
			if index == len(w)-1 {
				return true
			}
			(*flag)[i][j] = 1
			b := dfs(w, board, i+1, j, flag, index+1) ||
				dfs(w, board, i-1, j, flag, index+1) ||
				dfs(w, board, i, j+1, flag, index+1) ||
				dfs(w, board, i, j-1, flag, index+1)
			if b {
				return true
			}
			(*flag)[i][j] = 0
			return false
		} else {
			return false
		}
	}
}

/**
最大回文
*/
func test(str string) string {
	ss := []byte(str)
	var s = 0
	var e = 1
	var m = 1
	ret := make([]byte, 0)
	for s < len(str) {
		e = m
		tmp := make([]byte, 0)
		for e <= len(str) {
			if j(ss[s:e]) {
				if (e - s) > len(tmp) {
					tmp = ss[s:e]
					m = e
				}
			}
			e++
		}
		if len(tmp) > len(ret) {
			ret = tmp
		}
		s++
	}
	return string(ret)
}

func j(b []byte) bool {
	if len(b)%2 == 0 {
		for i := 0; i < len(b)/2; i++ {
			if b[i] != b[len(b)-i-1] {
				return false
			}
		}
	} else {
		for i := 0; i < (len(b)-1)/2; i++ {
			if b[i] != b[len(b)-i-1] {
				return false
			}
		}
	}
	return true
}

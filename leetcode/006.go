package leetcode

import (
	"math"
)

func Convert(s string, numRows int) string {
	if s == "" || numRows == 1 {
		return s
	}
	c := []byte(s)
	res := make([]byte, 0, len(s))
	rN := numRows
	bN := int(math.Ceil(float64(len(c)-rN) / float64(2*rN-2)))
	cN := bN*(rN-1) + 1
	i := 0
	for i < rN {
		j := 0
		for j < cN {
			if j-i > 0 && (j-i)%(rN-1) != 0 && 2*j-i < len(s) {
				res = append(res, c[2*j-i])
			}
			if 2*j+i < len(s) {
				res = append(res, c[2*j+i])
			}
			j += rN - 1
		}
		i++
	}
	return string(res)
}

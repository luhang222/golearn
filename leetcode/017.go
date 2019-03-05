package leetcode

func LetterCombinations(digits string) []string {
	var res []string
	if len(digits) == 0 {
		return res
	}
	mp := map[string]string{
		"2": "abc0",
		"3": "def0",
		"4": "ghi0",
		"5": "jkl0",
		"6": "mno0",
		"7": "pqrs",
		"8": "tuv0",
		"9": "wxyz",
	}
	route := make([][]byte, 0, len(digits))
	for _, a := range digits {
		route = append(route, []byte(mp[string(a)]))
	}
	for i := 0; i < 4; i++ {
		var path string
		walk(i, 0, route, path, &res)
	}
	return res
}

func walk(i int, height int, route [][]byte, path string, res *[]string) {
	if string(route[height][i]) == "0" {
		return
	}
	path += string(route[height][i])
	if height == len(route)-1 {
		*res = append(*res, path)
		return
	}
	switch i {
	case 0:
		walk(i+1, height+1, route, path, res)
		walk(i+2, height+1, route, path, res)
		walk(i+3, height+1, route, path, res)
	case 1:
		walk(i-1, height+1, route, path, res)
		walk(i+1, height+1, route, path, res)
		walk(i+2, height+1, route, path, res)
	case 2:
		walk(i-1, height+1, route, path, res)
		walk(i-2, height+1, route, path, res)
		walk(i+1, height+1, route, path, res)
	case 3:
		walk(i-1, height+1, route, path, res)
		walk(i-2, height+1, route, path, res)
		walk(i-3, height+1, route, path, res)
	}
	walk(i, height+1, route, path, res)
}

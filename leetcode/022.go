package leetcode

func GenerateParenthesis(n int) []string {
	res := make([]string, 0, 2*n)
	if n == 0 {
		return []string{}
	}
	if n == 1 {
		return []string{"()"}
	}
	add("(", n, n, "", &res)
	return res
}

func add(op string, remainLeft int, remainRight int, path string, res *[]string) {
	if op == "(" {
		remainLeft -= 1
	} else {
		remainRight -= 1
	}
	if remainLeft > remainRight {
		return
	}
	path += op
	if remainLeft > 0 {
		add("(", remainLeft, remainRight, path, res)
	}
	if remainRight > 0 {
		add(")", remainLeft, remainRight, path, res)
	}
	if remainLeft == 0 && remainRight == 0 {
		*res = append(*res, path)
		return
	}
}

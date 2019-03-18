package leetcode

/**
输入: candidates = [2,3,6,7], target = 7,
所求解集为:
[
  [7],
  [2,2,3]
]
*/
func CombinationSum(candidates []int, target int) [][]int {
	solves := make([][]int, 0)
	dfs(candidates, target, &solves, []int{})
	return solves
}

func dfs(candidates []int, target int, solves *[][]int, solve []int) bool {
	f := false
	for i := 0; i < len(candidates); i++ {
		j := 1
		for {
			t := target - candidates[i]*j
			for k := 0; k < j; k++ {
				solve = append(solve, candidates[i])
			}
			if t == 0 {
				f = true
				solution := make([]int, len(solve))
				copy(solution, solve)
				*solves = append(*solves, solution)
			} else if t > 0 {
				f = dfs(candidates[i+1:], t, solves, solve)
			} else {
				solve = solve[:len(solve)-j]
				break
			}
			solve = solve[:len(solve)-j]
			j++
		}
	}
	return f
}

//func CombinationSum2(candidates []int, target int) [][]int {
//	res := [][]int{}
//	cur := []int{}
//	combination2(candidates, target, cur, &res, 0)
//	return res
//}
//
//func combination2(candidates []int, target int, cur []int, res *[][]int, begin int) {
//	if target == 0 {
//		temp := make([]int, len(cur))
//		copy(temp, cur)
//		*res = append(*res, temp)
//		return
//	} else if target < 0 {
//		return
//	}
//	for i := begin; i < len(candidates); i++ {
//		cur = append(cur, candidates[i])
//		combination2(candidates, target-candidates[i], cur, res, i)
//		cur = cur[:len(cur)-1]
//	}
//}

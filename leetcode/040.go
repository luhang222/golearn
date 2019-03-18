package leetcode

import (
	"sort"
)

func CombinationSum2(candidates []int, target int) [][]int {
	solves := make([][]int, 0)
	mp := make(map[int]int, 0)
	for i := 0; i < len(candidates); i++ {
		mp[candidates[i]]++
	}
	sort.Ints(candidates)
	dfs2(candidates, target, &solves, []int{})
	return solves
}

func dfs2(candidates []int, target int, solves *[][]int, solve []int) bool {
	//fmt.Println(candidates)
	f := false
	for i := 0; i < len(candidates); i++ {
		t := target - candidates[i]
		solve = append(solve, candidates[i])
		if t == 0 {
			if i > 0 && candidates[i-1] == candidates[i] {
				solve = solve[:len(solve)-1]
				continue
			}
			f = true
			solution := make([]int, len(solve))
			copy(solution, solve)
			*solves = append(*solves, solution)
		} else if t > 0 {
			if i > 0 && candidates[i-1] == candidates[i] {
				solve = solve[:len(solve)-1]
				continue
			}
			f = dfs2(candidates[i+1:], t, solves, solve)
		}
		solve = solve[:len(solve)-1]
	}
	return f
}

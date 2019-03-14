package leetcode

/**
可以优化，两次二分，先找左边再找右边即可
*/
func SearchRange(nums []int, target int) []int {
	l := len(nums)
	i, j := 0, l-1
	for i <= j {
		c := (i + j) / 2
		if target > nums[c] {
			i = c + 1
		} else if target < nums[c] {
			j = c - 1
		} else {
			ii, jj, ci, cj := i, j, c, c
			if nums[ii] != target {
				for ii < ci {
					cc := nums[(ii+ci)/2]
					if cc != target {
						if ii == (ii+ci)/2 {
							ii = ci
							break
						} else {
							ii = (ii + ci) / 2
						}
					} else {
						ci = (ii + ci) / 2
					}
				}
			}
			if nums[jj] != target {
				for jj > cj {
					cc := nums[(cj+jj)/2]
					if cc != target {
						jj = (cj + jj) / 2
					} else {
						if cj == (cj+jj)/2 {
							jj = cj
							break
						} else {
							cj = (cj + jj) / 2
						}
					}
				}
			}
			return []int{ii, jj}
		}
	}
	return []int{-1, -1}
}

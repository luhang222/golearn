package leetcode

import "strings"

func IsMatch(s string, p string) bool {
	index := strings.Index(p, "*")
	if index == -1 {
		if len(p) != len(s) {
			return false
		}
		for i := range s {
			if s[i] != p[i] && p[i] != 46 {
				return false
			}
		}
		return true
	} else {
		subp := p[:index]
		if len(subp) == 0 {
			return false
		}
		for i := 0; i < len(subp)-1; i++ {
			if len(s) < i+1 {
				return false
			}
			if s[i] != subp[i] && subp[i] != 46 {
				return false
			}
		}
		remain := s[len(subp)-1:]
		last := subp[len(subp)-1]
		//0
		if IsMatch(remain, p[index+1:]) {
			return true
		}
		for i := 0; i < len(remain); i++ {
			if remain[i] != last && last != 46 {
				return false
			} else {
				if IsMatch(remain[i+1:], p[index+1:]) {
					return true
				}
			}
		}
	}
	return false
}

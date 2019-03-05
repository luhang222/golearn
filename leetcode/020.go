package leetcode

func IsValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s) == 1 {
		return false
	}
	// ( 40 ) 41 { 123 } 125 [ 91 ] 93
	stk := make([]byte, 0, len(s))
	for _, b := range s {
		if b == 40 || b == 123 || b == 91 {
			stk = append(stk, byte(b))
		} else {
			if len(stk) == 0 {
				return false
			}
			switch stk[len(stk)-1] {
			case 40:
				if b != 41 {
					return false
				}
				stk = stk[:len(stk)-1]
			case 123:
				if b != 125 {
					return false
				}
				stk = stk[:len(stk)-1]
			case 91:
				if b != 93 {
					return false
				}
				stk = stk[:len(stk)-1]
			default:
				return false
			}
		}
	}
	return len(stk) == 0
}

package piscine

func AtoiBase(s string, base string) int {
	daibou := 0
	yade := 0
	for i, v := range base {
		if v == '-' || v == '+' {
			return 0
		}
		daibou++
		for j, v2 := range base {
			if v == v2 && i != j {
				return 0
			}
		}
	}
	for _, v := range s {
		if v == '-' || v == '+' {
			return 0
		}
		yade++
	}
	if daibou < 2 {
		return 0
	}
	runes := []rune(s)
	result := 0
	j := 0
	for i := yade - 1; i >= 0; i-- {
		r := -1
		for k, v := range base {
			if v == runes[i] {
				r = k
			}
		}
		if r == (-1) {
			return 0
		}
		if j == 0 {
			result += r
		} else {
			ss := 1
			for l := 0; l < j; l++ {
				ss *= daibou
			}
			result += ss * r
		}
		j++
	}
	return result
}

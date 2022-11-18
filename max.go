package piscine

func Max(a []int) int {
	max := a[0]
	if max == 0 {
		return 0
	}
	for i := range a {
		if a[i] > max {
			max = a[i]
		}
	}
	return max
}

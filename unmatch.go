package piscine

func Unmatch(a []int) int {
	pair := 0
	for i := range a {
		pair = i + 1
	}
	for i := 0; i < pair; i++ {
		for j := i + 1; j < pair; j++ {
			if a[i] == a[j] {
				a[i], a[j] = -1, -1
			}
		}
	}
	for _, i := range a {
		if i != -1 {
			return i
		}
	}

	return -1
}

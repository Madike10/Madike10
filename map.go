package piscine

func Map(f func(int) bool, a []int) []bool {
	yade := make([]bool, len(a))

	for i, j := range a {
		yade[i] = f(j)
	}
	return yade
}

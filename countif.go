package piscine

func CountIf(f func(string) bool, tab []string) int {
	yade := 0
	for _, a := range tab {
		if f(a) {
			yade++
		}
	}
	return yade
}

package piscine

func BasicAtoi(s string) int {
	yade := 0
	for _, n := range s {
		yade = yade*10 + int(n) - '0'
	}
	return yade
}

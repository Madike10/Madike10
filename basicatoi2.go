package piscine

func BasicAtoi2(s string) int {
	yade := 0
	for _, n := range s {
		if n < '0' || n > '9' {
			return 0
		}
		yade = yade*10 + int(n) - '0'
	}
	return yade
}

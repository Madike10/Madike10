package piscine

func ToLower(s string) string {
	dickss := []rune(s)
	for i := 0; i < len(dickss); i++ {
		if dickss[i] >= 'A' && dickss[i] <= 'Z' {
			dickss[i] += 32
		}
	}
	return string(dickss)
}

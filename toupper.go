package piscine

func ToUpper(s string) string {
	dickss := []rune(s)
	for i := 0; i < len(dickss); i++ {
		if dickss[i] >= 'a' && dickss[i] <= 'z' {
			dickss[i] -= 32
		}
	}
	return string(dickss)
}

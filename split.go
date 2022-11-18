package piscine

func Split(s, sep string) []string {
	var yade string
	IncLetter := 0
	tab := make([]string, 0)

	for i := 0; i < len(s); i++ {
		count := 0
		for j := 0; j < len(sep); j++ {
			if i+j < len(s) && sep[j] == s[i+j] {
				count += 1
			}
		}
		if count != len(sep) {
			yade += string(s[i])
			IncLetter += 1
		} else {
			tab = append(tab, yade)
			yade = ""
			i = i + len(sep) - 1
			IncLetter = 0
		}
		if i == len(s)-1 && IncLetter > 0 {
			tab = append(tab, yade)
		}
	}
	return tab
}

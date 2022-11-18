package piscine

func SplitWhiteSpaces(s string) []string {
	var yade string = ""
	var tab []string
	for i := 0; i < len(s); i++ {
		if s[i] > ' ' {
			yade = yade + string(s[i])
		} else if yade != "" {
			tab = append(tab, yade)
			yade = ""
		}
		if i == len(s)-1 {
			tab = append(tab, yade)
			yade = ""

		}
	}
	return tab
}

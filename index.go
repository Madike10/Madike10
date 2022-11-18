package piscine

func Index(s string, toFind string) int {
	index := -1
	a := []rune(s)
	b := []rune(toFind)

	for i := 0; i < len(s); i++ {
		if len(b)+i <= len(a) {
			if s[i:len(b)+i] == toFind {
				index = i
				break
			}
		}
	}
	return index
}

package piscine

func Join(strs []string, sep string) string {
	ymadike := ""
	for i := range strs {
		if i != 0 {
			ymadike = Concat(ymadike, sep)
		}
		ymadike = Concat(ymadike, strs[i])
	}
	return ymadike
}

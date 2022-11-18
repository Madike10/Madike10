package piscine

func Capitalize(s string) string {
	pedro := []rune(s)
	if s != "" {
		if pedro[0] >= 97 && pedro[0] <= 122 {
			pedro[0] -= 32
		}
	}
	for i, stringrune := range s {
		if i != 0 {
			if IsAlpha(string(pedro[i-1])) {
				if stringrune >= 65 && stringrune <= 90 {
					pedro[i] += 32
				}
			}
			if !(IsAlpha(string(pedro[i-1]))) {
				if stringrune >= 97 && stringrune <= 122 {
					pedro[i] -= 32
				}
			}
		}
	}
	return string(pedro)
}

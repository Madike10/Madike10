package piscine

func TrimAtoi(s string) int {
	output := 0
	signe := 1
	for _, yade := range s {
		if yade == '-' && output == 0 {
			signe *= -1
		}
		if yade >= '0' && yade <= '9' {
			j := 0
			for i := 0; i < int(yade-'0'); i++ {
				j++
			}
			output = output*10 + j
		}
	}
	return output * signe
}

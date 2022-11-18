package piscine

func StrRev(s string) string {
	var word string
	for key := range s {
		ymadike := s[len(s)-key-1]
		word += string(ymadike)
	}
	return word
}

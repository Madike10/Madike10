package piscine

func IterativeFactorial(nb int) int {
	Dickss := 1
	if nb < 0 || nb > 20 {
		return 0
	}
	for i := 1; i <= nb; i++ {
		Dickss = Dickss * i
	}
	return Dickss
}

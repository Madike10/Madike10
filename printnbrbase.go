package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	counter := 0
	for _, r1 := range base {
		yade := 0
		if r1 == '+' || r1 == '-' {
			z01.PrintRune('N')
			z01.PrintRune('V')
			return
		}
		for _, r2 := range base {
			if r1 == r2 {
				yade++
			}
			if yade == 2 {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
		counter++
	}
	if counter < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	output := ""
	ra := []rune(base)
	if nbr < 0 {
		z01.PrintRune('-')
	}
	for ; nbr != 0; nbr /= counter {
		index := nbr % counter
		if index < 0 {
			index = -index
		}
		output = string(ra[index]) + output
	}
	for _, r := range output {
		z01.PrintRune(r)
	}
}

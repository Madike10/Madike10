package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, str := range a {
		for _, c := range []rune(str) {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
	}
}

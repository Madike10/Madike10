package piscine

import "github.com/01-edu/z01"

func PrintNum(a int) {
	i := '0'
	if a == 0 {
		z01.PrintRune('0')
		return
	}
	for j := 1; j <= a%10; j++ {
		i++
	}
	for j := -1; j >= a%10; j-- {
		i++
	}
	if a/10 != 0 {
		PrintNum(a / 10)
	}
	z01.PrintRune(i)
}

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	PrintNum(n)
}

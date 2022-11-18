package piscine

func AppendRange(min, max int) []int {
	var tab []int

	for i := min; i < max; i++ {
		if min >= max {
			return nil
		}

		tab = append(tab, i)
	}
	return tab
}

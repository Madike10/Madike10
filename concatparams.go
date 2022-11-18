package piscine

func ConcatParams(args []string) string {
	tab := args[0]
	for i := 1; i < len(args); i++ {
		tab += "\n" + args[i]
	}
	return tab
}

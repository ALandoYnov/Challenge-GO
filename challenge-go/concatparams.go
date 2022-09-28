package isnegative

func ConcatParams(args []string) string {
	var a string
	for i := 0; i < len(args); i++ {
		if len(args)-1 > i {
			a += args[i]
			a += "\n"
		} else {
			a += args[i]
		}
	}
	return a
}

package isnegative

func IterativeFactorial(nb int) int {
	resultat := 1

	if nb > 20 || nb < 0 {
		return 0
	} else if nb != 0 {
		for i := 1; i <= nb; i++ {
			resultat = resultat * i
		}
		return resultat
	} else {
		return 1
	}
}

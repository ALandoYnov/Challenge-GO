package isnegative

func UltimateDivMod(div *int, mod *int) {
	y := *div
	*div = *div / *mod
	*mod = y % *mod
}

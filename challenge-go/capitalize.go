package isnegative

func Capitalize(s string) string {
	var resultat string
	for i := range s {
		if IsAlpha(string(s[i])) {
			if IsUpper(string(s[i])) {
				if i == 0 {
					resultat = resultat + string(s[i])
					continue
				} else if IsUpper(string(s[i-1])) {
					resultat = resultat + ToLower(string(s[i]))
					continue
				} else if IsLower(string(s[i-1])) {
					resultat = resultat + ToLower(string(s[i]))
					continue
				} else if s[i-1] >= 48 && s[i-1] <= 57 {
					resultat = resultat + ToLower(string(s[i]))
					continue
				} else {
					resultat = resultat + string(s[i])
					continue
				}
			} else if IsLower(string(s[i])) {
				if i == 0 {
					resultat = resultat + ToUpper(string(s[i]))
					continue
				} else if !(IsAlpha(string(s[i-1]))) {
					resultat = resultat + ToUpper(string(s[i]))
					continue
				} else {
					resultat = resultat + string(s[i])
					continue
				}
			}
			resultat = resultat + string(s[i])
			continue
		}
		resultat = resultat + string(s[i])
	}
	return resultat
}

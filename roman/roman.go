package roman

// NToI erwartet eine Zahl und liefert die entsprechende Anzahl an I als String.
func NToI(n int) string {
	number := ""
	for ; n > 0; n-- {
		number = number + "I"
	}
	return number
}

// NToX erwartet eine Zahl und liefert die entsprechende Anzahl an X als String.
func NToX(n int) string {
	number := ""
	n = n / 10
	for ; n > 0; n-- {
		number = number + "X"
	}
	return number
}

// RomanBelow10 erwartet eine Zahl 1 <= n <= 10 und liefert die römische Schreibweise für n als String.
func RomanBelow10(n int) string {
	number := ""
	switch n {
	case 1:
		number = number + "I"
	case 2:
		number = number + "II"
	case 3:
		number = number + "III"
	case 4:
		number = number + "IV"
	case 5:
		number = number + "V"
	case 6:
		number = number + "VI"
	case 7:
		number = number + "VII"
	case 8:
		number = number + "VIII"
	case 9:
		number = number + "IX"
	case 10:
		number = number + "X"
	}

	return number
}

// RomanBelow100 erwartet eine Zahl 1 <= n <= 100 und liefert die römische Schreibweise für n als String.
func RomanBelow100(n int) string {
	number := ""
	switch n {
	case 50:
		number = "L"
	case 100:
		number = "M"
	}
	return number
}

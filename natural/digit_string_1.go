package natural

// DigitString1 erwartet eine Ziffer als int.
// Die Funktion gibt den zugehörigen String zurück, wie er üblicherweise
// an der Einer-Stelle einer Zahl >= 21 vorkommen würde.
// Außerdem wird bei Ziffern != 0 das Wort "und" angehängt.
//
// Beispiel: Für 3 soll der String "dreiund" geliefert werden.
//
// Anmerkung:
// Dies ist eine Hilfsfunktion, die genutzt werden soll,
// um den Gesamt-String einer Zahl zusammenzusetzen.
// Diese Funktion muss nur für den Normalfall (Zahlen >= 21) funktionieren.
func DigitString1(digit int) string {
	endnumber := ""
	number := ""
	if digit == 0 {
		return number
	}
	switch digit {
	case 1:
		number = "ein"
	case 2:
		number = "zwei"
	case 3:
		number = "drei"
	case 4:
		number = "vier"
	case 5:
		number = "fünf"
	case 6:
		number = "sechs"
	case 7:
		number = "sieben"
	case 8:
		number = "acht"
	case 9:
		number = "neun"
	}

	endnumber = number + "und"
	return endnumber
}

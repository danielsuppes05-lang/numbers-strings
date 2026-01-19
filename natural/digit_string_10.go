package natural

// DigitString10 erwartet eine Ziffer als int.
// Die Funktion gibt den zugehörigen String zurück, wie er üblicherweise
// an der Zehner-Stelle einer Zahl >= 21 vorkommen würde.
//
// Beispiel: Für 3 soll der String "dreißig" geliefert werden.
//
// Anmerkung:
// Dies ist eine Hilfsfunktion, die genutzt werden soll,
// um den Gesamt-String einer Zahl zusammenzusetzen.
func DigitString10(digit int) string {
	endnumber := ""
	number := ""
	switch digit {
	case 0:
		return number
	case 1:
		return number
	case 2:
		endnumber = endnumber + "zwan"
		return endnumber
	case 3:
		endnumber = endnumber + "dreiß"
		return endnumber
	}
	s := DigitString1(digit)
	s = s[:len(s)-3]
	endnumber = s + "zig"
	return endnumber
}

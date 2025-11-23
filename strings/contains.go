package strings

// Erwartet einen String s und einen Buchstaben c.
// Prüft, ob c in s vorkommt.
func Contains(s string, c byte) bool {
	for i := 0; i < len(s); i++ { //i startet bei 0, gfeht die Länge des strings s in einerschritten durch, solange i kleiner ist als Länge s
		if c == s[i] { //wenn c ein Element vom String s ist, return true
			return true
		}
	}
	return false

}

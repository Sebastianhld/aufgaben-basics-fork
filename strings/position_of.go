package strings

// Erwartet einen String s und einen Buchstaben c.
// Liefert die Position, an der c in s vorkommt.
// Liefert die Länge von s, falls c nicht in s vorkommt.
// Kommt c mehrfach vor, soll die erste Position geliefert werden.
func PositionOf(s string, c byte) int {
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return i
		}
	}

	return len(s)
}
